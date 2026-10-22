package services

import (
	"errors"
	"gin-layered-api/models"
	"gin-layered-api/repositories"
	"gin-layered-api/utils"
)

type AuthService interface {
	Login(req models.LoginRequest) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(req models.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil || user == nil || user.Password != req.Password {
		return "", errors.New("invalid username or password")
	}

	return utils.GenerateToken(user.Username, user.Role)
}
