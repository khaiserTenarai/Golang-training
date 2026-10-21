package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	PasswordHash []byte
	Role         string
}

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
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
				Username: adminUsername, PasswordHash: adminHash, Role: "admin",
			},
			userUsername: {
				Username: userUsername, PasswordHash: userHash, Role: "user",
			},
		},
	}, nil
}

// Login verifies credentials and returns a signed JWT.
func (a *AuthService) Login(username, password string) (string, error) {
	u, ok := a.users[username]
	if !ok || bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) != nil {
		return "", errors.New("invalid username or password")
	}

	now := time.Now()
	claims := Claims{
		Role: u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(a.expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}

// Parse validates the signature, algorithm, and expiration of a bearer token.
func (a *AuthService) ParseAuthorization(header string) (*Claims, error) {
	const prefix = "Bearer "
	if len(header) <= len(prefix) || header[:len(prefix)] != prefix {
		return nil, errors.New("missing bearer token")
	}

	raw := header[len(prefix):]
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		raw,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected JWT signing algorithm")
			}
			return a.secret, nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
	)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	if claims.Subject == "" || (claims.Role != "admin" && claims.Role != "user") {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	// Confirm the subject and role match a currently configured user.
	u, ok := a.users[claims.Subject]
	if !ok || u.Role != claims.Role {
		return nil, errors.New("unknown user or invalid role")
	}

	return claims, nil
}
