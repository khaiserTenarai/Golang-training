package service

import (
	"employee-microservices/auth/repository"
	"errors"
)

// AuthService contains login rules. Task 6. This demo intentionally does not issue JWTs.
type AuthService struct{ repo *repository.UserRepository }

func NewAuthService(r *repository.UserRepository) *AuthService { return &AuthService{repo: r} }
func (s *AuthService) Login(username, password string) error {
	if !s.repo.IsValidDemoUser(username, password) {
		return errors.New("invalid demo credentials")
	}
	return nil
}
