package service

import (
	"context"
	"strings"
	"time"

	"secure-banking-api/config"
	"secure-banking-api/model"
	"secure-banking-api/repository"
	"secure-banking-api/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (string, error)
}

type authService struct {
	users repository.UserRepository
	cfg   *config.Config
}

func NewAuthService(users repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{users: users, cfg: cfg}
}

func (s *authService) Register(ctx context.Context, req model.RegisterRequest) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	role := model.RoleUser
	if s.cfg.AdminEmail != "" && email == s.cfg.AdminEmail {
		role = model.RoleAdmin
	}

	u := &model.User{Name: strings.TrimSpace(req.Name), Email: email, PasswordHash: string(hash), Role: role}
	if err := s.users.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *authService) Login(ctx context.Context, req model.LoginRequest) (string, error) {
	u, err := s.users.FindByEmail(ctx, strings.ToLower(strings.TrimSpace(req.Email)))
	if err != nil {
		if err == model.ErrNotFound {
			return "", model.ErrInvalidCredentials 
		}
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)) != nil {
		return "", model.ErrInvalidCredentials
	}
	return util.GenerateToken(s.cfg.JWTSecret, u.ID, u.Role, time.Duration(s.cfg.JWTExpiryMinutes)*time.Minute)
}
