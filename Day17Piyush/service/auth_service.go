// Tasks 6, 7: Authentication and Authorization service

package service

import (
	"fmt"
	"sync"
	"time"

	"graphql_employee/models"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("day17-graphql-secret-key")

type AuthService struct {
	mu    sync.RWMutex
	users map[string]models.User
	next  int
}

func NewAuthService() *AuthService {
	s := &AuthService{users: make(map[string]models.User), next: 1}
	seed := []models.User{
		{Username: "admin", Password: "admin123", Role: "admin"},
		{Username: "manager", Password: "manager123", Role: "manager"},
		{Username: "piyush", Password: "piyush123", Role: "user"},
	}
	for _, u := range seed {
		u.ID = s.next
		s.users[u.Username] = u
		s.next++
	}
	return s
}

func (s *AuthService) Login(username, password string) (string, models.User, error) {
	s.mu.RLock()
	user, ok := s.users[username]
	s.mu.RUnlock()
	if !ok || user.Password != password {
		return "", models.User{}, fmt.Errorf("invalid credentials")
	}
	token, err := GenerateJWT(user.ID, user.Username, user.Role)
	return token, user, err
}

func (s *AuthService) Register(username, password, role string) (string, models.User, error) {
	if len(username) < 3 {
		return "", models.User{}, fmt.Errorf("username must be at least 3 characters")
	}
	if len(password) < 6 {
		return "", models.User{}, fmt.Errorf("password must be at least 6 characters")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[username]; exists {
		return "", models.User{}, fmt.Errorf("username taken")
	}
	if role == "" {
		role = "user"
	}
	user := models.User{ID: s.next, Username: username, Password: password, Role: role}
	s.next++
	s.users[username] = user
	token, _ := GenerateJWT(user.ID, user.Username, user.Role)
	return token, user, nil
}

func GenerateJWT(id int, username, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JWTSecret)
}

func ParseJWT(tokenStr string) (int, string, string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, "", "", fmt.Errorf("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	return int(claims["user_id"].(float64)), claims["username"].(string), claims["role"].(string), nil
}
