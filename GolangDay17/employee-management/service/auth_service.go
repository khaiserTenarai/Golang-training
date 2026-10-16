package service

import (
	"employee-management/auth"
	"employee-management/model"
	"employee-management/repository"
	"strings"
)


type LoginResult struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type AuthService interface {
	Login(username, password string) (*LoginResult, error)
	SeedDefaultUsers() error
}
type authService struct {
	users  repository.UserRepository
	tokens *auth.TokenService
}

func NewAuthService(u repository.UserRepository, t *auth.TokenService) AuthService {
	return &authService{u, t}
}


func (s *authService) Login(username, password string) (*LoginResult, error) {
	u, err := s.users.GetByUsername(strings.TrimSpace(username))
	if err != nil {
		if IsNotFound(err) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	if !auth.CheckPassword(u.PasswordHash, password) {
		return nil, ErrInvalidCredentials
	}
	token, err := s.tokens.Generate(u.ID, u.Username, u.Role)
	if err != nil {
		return nil, err
	}
	return &LoginResult{Token: token, Username: u.Username, Role: u.Role}, nil
}


func (s *authService) SeedDefaultUsers() error {
	n, err := s.users.Count()
	if err != nil || n > 0 {
		return err
	}
	defaults := []struct{ name, pass, role string }{
		{"admin", "admin123", auth.RoleAdmin},
		{"user", "user123", auth.RoleUser},
	}
	for _, d := range defaults {
		hash, err := auth.HashPassword(d.pass)
		if err != nil {
			return err
		}
		if err := s.users.Create(&model.User{Username: d.name, PasswordHash: hash, Role: d.role}); err != nil {
			return err
		}
	}
	return nil
}
