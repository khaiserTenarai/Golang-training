package service

import (
	"errors"
	"fmt"
	"time"

	"order-management/config"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	secret []byte
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthService() *AuthService {
	return &AuthService{
		secret: []byte(config.Get(
			"JWT_SECRET",
			"replace-with-a-long-random-secret-at-least-32-characters",
		)),
	}
}

func (s *AuthService) Login(username, password string) (string, string, error) {
	var role string

	switch {
	case username == "admin" && password == "admin@123":
		role = "admin"
	case username == "user" && password == "user@123":
		role = "user"
	default:
		return "", "", errors.New("invalid username or password")
	}

	now := time.Now()
	claims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Hour)),
			Issuer:    "order-management",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.secret)
	if err != nil {
		return "", "", fmt.Errorf("sign JWT: %w", err)
	}

	return signedToken, role, nil
}
