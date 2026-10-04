package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"task15_employee_management_capstone/models"
	"task15_employee_management_capstone/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user models.User) (models.User, error)
	Login(username, password string) (string, models.User, error)
	ValidateToken(token string) (int, string, error)
	Logout(token string) error
	GetUserByID(id int) (models.User, error)
}

type AuthServiceImpl struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Register(user models.User) (models.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return models.User{}, errors.New("username, email, and password are required")
	}
	if len(user.Password) < 6 {
		return models.User{}, errors.New("password must be at least 6 characters")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("failed to hash password")
	}
	user.PasswordHash = string(hashed)
	if user.Role == "" {
		user.Role = "user"
	}
	return s.repo.Create(user)
}

func (s *AuthServiceImpl) Login(username, password string) (string, models.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", models.User{}, errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", models.User{}, errors.New("invalid username or password")
	}
	token, _ := generateToken()
	s.repo.CreateSession(user.ID, token, time.Now().Add(24*time.Hour))
	user.PasswordHash = ""
	return token, user, nil
}

func (s *AuthServiceImpl) ValidateToken(token string) (int, string, error) {
	return s.repo.GetSession(token)
}

func (s *AuthServiceImpl) Logout(token string) error {
	return s.repo.DeleteSession(token)
}

func (s *AuthServiceImpl) GetUserByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes), nil
}
