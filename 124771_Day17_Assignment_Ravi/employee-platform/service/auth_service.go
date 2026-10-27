package service

import (
	"errors"
	"strings"

	"employee-platform/model"
	"employee-platform/repository"
	"employee-platform/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	users  *repository.UserRepository
	secret string
	hours  int
}

func NewAuthService(users *repository.UserRepository, secret string, hours int) *AuthService {
	return &AuthService{users: users, secret: secret, hours: hours}
}
func (s *AuthService) Register(name, email, password string) (*model.User, error) {
	name, email = strings.TrimSpace(name), strings.ToLower(strings.TrimSpace(email))
	if name == "" || email == "" || len(password) < 8 {
		return nil, errors.New("name/email required and password must be at least 8 characters")
	}
	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	// Task 7: public registration always assigns employee; clients cannot self-assign admin.
	user := &model.User{Name: name, Email: email, PasswordHash: hash, Role: "employee"}
	if err := s.users.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
func (s *AuthService) Login(email, password string) (*model.User, string, error) {
	user, err := s.users.ByEmail(strings.ToLower(strings.TrimSpace(email)))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("invalid email or password")
		}
		return nil, "", err
	}
	if !utils.CheckPassword(user.PasswordHash, password) {
		return nil, "", errors.New("invalid email or password")
	}
	token, err := utils.GenerateToken(user.ID, user.Role, s.secret, s.hours)
	return user, token, err
}
func (s *AuthService) SeedAdmin(email, password, name string) error {
	_, err := s.users.ByEmail(strings.ToLower(email))
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return s.users.Create(&model.User{Name: name, Email: strings.ToLower(email), PasswordHash: hash, Role: "admin"})
}
