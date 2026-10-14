package main

import (
	"errors"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

const tokenLifetime = time.Hour

var now = time.Now

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func initJWT() {
	secret := os.Getenv("JWT_SECRET")

	if len(secret) < 32 {
		panic("JWT_SECRET must contain at least 32 characters")
	}

	jwtSecret = []byte(secret)
}

func generateToken(user User) (string, error) {
	currentTime := now()

	claims := Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(tokenLifetime)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func checkPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	) == nil
}

func validEmail(email string) bool {
	email = strings.TrimSpace(email)
	parts := strings.Split(email, "@")

	return len(parts) == 2 &&
		parts[0] != "" &&
		strings.Contains(parts[1], ".") &&
		!strings.ContainsAny(email, " \t\r\n")
}

func validPassword(password string) bool {
	return utf8.RuneCountInString(password) >= 8 &&
		len(password) <= 72
}

func validateTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return errors.New("title is required")
	}

	if utf8.RuneCountInString(title) > 100 {
		return errors.New("title must be at most 100 characters")
	}

	return nil
}
