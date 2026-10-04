package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"task13_authentication_system/models"
	"task13_authentication_system/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(user models.User) (models.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return models.User{}, errors.New("username, email, and password are required")
	}
	if len(user.Password) < 6 {
		return models.User{}, errors.New("password must be at least 6 characters")
	}

	// Hash the password using bcrypt
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("failed to hash password")
	}
	user.PasswordHash = string(hashedBytes)

	if user.Role == "" {
		user.Role = "user"
	}

	return s.Repo.Create(user)
}

func (s *AuthService) Login(username, password string) (string, models.User, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		return "", models.User{}, errors.New("invalid username or password")
	}

	// Compare password with bcrypt hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", models.User{}, errors.New("invalid username or password")
	}

	// Generate session token
	token, err := generateToken()
	if err != nil {
		return "", models.User{}, errors.New("failed to generate session token")
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	if err := s.Repo.CreateSession(user.ID, token, expiresAt); err != nil {
		return "", models.User{}, errors.New("failed to create session")
	}

	user.PasswordHash = ""
	return token, user, nil
}

func (s *AuthService) ValidateToken(token string) (int, string, error) {
	return s.Repo.GetSession(token)
}

func (s *AuthService) Logout(token string) error {
	return s.Repo.DeleteSession(token)
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
