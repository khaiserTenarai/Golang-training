package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	// Import JWT implementation.
	"github.com/golang-jwt/jwt/v5"
)

// User represents a configured application user.
type User struct {
	Username     string
	PasswordHash string
	Role         string
}

// AuthService validates credentials and JWTs.
type AuthService struct {
	Secret []byte
	Expiry time.Duration
	Admin  User
	Normal User
}

func New(secret string, expiryMinutes int, adminUsername, adminHash, userUsername, userHash string) *AuthService {
	return &AuthService{Secret: []byte(secret), Expiry: time.Duration(expiryMinutes) * time.Minute, Admin: User{adminUsername, adminHash, "admin"}, Normal: User{userUsername, userHash, "user"}}
}

func (a *AuthService) Login(username, password string) (string, error) {
	u := a.findUser(username)
	if u == nil || !compareHash(password, u.PasswordHash) {
		return "", errors.New("invalid username or password")
	}
	claims := jwt.MapClaims{"sub": u.Username, "role": u.Role, "exp": time.Now().Add(a.Expiry).Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.Secret)
}

func (a *AuthService) Parse(header string) (string, string, error) {
	parts := strings.Fields(header)
	//Validate the header format
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", "", errors.New("invalid authorization header")
	}

	token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return a.Secret, nil
	})
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid or expired token")
	}
	//Extract token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token claims")
	}
	user, _ := claims["sub"].(string)
	role, _ := claims["role"].(string)
	return user, role, nil
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
func compareHash(password, expected string) bool {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:]) == expected
}
