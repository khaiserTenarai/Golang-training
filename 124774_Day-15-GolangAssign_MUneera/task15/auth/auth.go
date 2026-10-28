package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username     string
	PasswordHash string
	Role         string
}

type AuthService struct {
	Secret []byte
	Expiry time.Duration
	Admin  User
	Normal User
}

func New(
	secret string,
	expiryMinutes int,
	adminUsername, adminHash string,
	userUsername, userHash string,
) *AuthService {
	return &AuthService{
		Secret: []byte(secret),
		Expiry: time.Duration(expiryMinutes) * time.Minute,
		Admin: User{
			Username:     adminUsername,
			PasswordHash: adminHash,
			Role:         "admin",
		},
		Normal: User{
			Username:     userUsername,
			PasswordHash: userHash,
			Role:         "user",
		},
	}
}

// HashPassword is for this training project's demo credentials.
// Production applications should use bcrypt or Argon2 for passwords.
func HashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}

func compareHash(password, expected string) bool {
	actual := HashPassword(password)
	return subtle.ConstantTimeCompare(
		[]byte(actual),
		[]byte(expected),
	) == 1
}

func (a *AuthService) Login(username, password string) (string, error) {
	user := a.findUser(username)

	if user == nil || !compareHash(password, user.PasswordHash) {
		return "", errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"sub":  user.Username,
		"role": user.Role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(a.Expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.Secret)
}

func (a *AuthService) Parse(header string) (string, string, error) {
	parts := strings.Fields(header)

	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", "", errors.New("invalid authorization header")
	}

	token, err := jwt.Parse(parts[1], func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return a.Secret, nil
	})

	if err != nil || token == nil || !token.Valid {
		return "", "", errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	username, userOK := claims["sub"].(string)
	role, roleOK := claims["role"].(string)

	if !userOK || !roleOK || username == "" || role == "" {
		return "", "", errors.New("invalid token claims")
	}

	return username, role, nil
}

func (a *AuthService) findUser(username string) *User {
	if username == a.Admin.Username {
		return &a.Admin
	}

	if username == a.Normal.Username {
		return &a.Normal
	}

	return nil
}
