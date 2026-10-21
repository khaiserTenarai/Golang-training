package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type User struct {
	Username     string
	PasswordHash []byte
	Role         string
}

type AuthService struct {
	secret []byte
	expiry time.Duration
	users  map[string]User
}

func New(
	secret string,
	expiryMinutes int,
	adminUsername, adminPassword string,
	userUsername, userPassword string,
) (*AuthService, error) {
	if len(secret) < 32 || expiryMinutes <= 0 {
		return nil, errors.New("invalid JWT configuration")
	}

	// Store password hashes, not plaintext passwords.
	adminHash, err := bcrypt.GenerateFromPassword(
		[]byte(adminPassword), bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	userHash, err := bcrypt.GenerateFromPassword(
		[]byte(userPassword), bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	return &AuthService{
		secret: []byte(secret),
		expiry: time.Duration(expiryMinutes) * time.Minute,
		users: map[string]User{
			adminUsername: {
				Username:     adminUsername,
				PasswordHash: adminHash,
				Role:         "admin",
			},
			userUsername: {
				Username:     userUsername,
				PasswordHash: userHash,
				Role:         "user",
			},
		},
	}, nil
}

func (a *AuthService) Login(username, password string) (string, error) {
	u, ok := a.users[username]

	if !ok || bcrypt.CompareHashAndPassword(
		u.PasswordHash, []byte(password),
	) != nil {
		return "", errors.New("invalid username or password")
	}

	now := time.Now()

	claims := Claims{
		Role: u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(a.expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}

func (a *AuthService) ParseToken(raw string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		raw,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing algorithm")
			}
			return a.secret, nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
	)

	if err != nil || token == nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	// Check the user's role against the configured users.
	u, ok := a.users[claims.Subject]
	if !ok || u.Role != claims.Role {
		return nil, errors.New("invalid user or role")
	}

	return claims, nil
}
