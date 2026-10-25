package service

import (
	"errors"
	"strings"
	"time"

	"employee-management-microservices/services/auth/model"
	"employee-management-microservices/services/auth/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type AuthService struct {
	repo      *repository.Repository
	jwtSecret []byte
}

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func New(repo *repository.Repository, secret string) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: []byte(secret),
	}
}

// Register creates an account with a hashed password.
func (s *AuthService) Register(username, password, role string) error {
	username = strings.TrimSpace(username)

	if username == "" || len(password) < 8 {
		return errors.New("username is required and password must have at least 8 characters")
	}

	role = strings.ToUpper(strings.TrimSpace(role))
	if role != "ADMIN" && role != "USER" {
		return errors.New("role must be ADMIN or USER")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username:     username,
		PasswordHash: string(hash),
		Role:         role,
		Active:       true,
	}

	return s.repo.Create(user)
}

// Login verifies credentials and returns a signed JWT.
func (s *AuthService) Login(username, password string) (string, *model.User, error) {
	user, err := s.repo.FindByUsername(strings.TrimSpace(username))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, ErrInvalidCredentials
		}
		return "", nil, err
	}

	if !user.Active ||
		bcrypt.CompareHashAndPassword(
			[]byte(user.PasswordHash), []byte(password),
		) != nil {
		return "", nil, ErrInvalidCredentials
	}

	now := time.Now()
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return signed, user, nil
}

// ValidateToken verifies the signature and expiry.
func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}
			return s.jwtSecret, nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)
	if err != nil || !token.Valid {
		return nil, ErrInvalidCredentials
	}

	return claims, nil
}
