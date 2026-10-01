package service

import (
	"context"
	"errors"
	"example.com/q13-authentication/model"
	"example.com/q13-authentication/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct{ repo repository.UserRepository }

func NewAuthService(r repository.UserRepository) AuthService { return &AuthServiceImpl{repo: r} }
func (s *AuthServiceImpl) Register(c context.Context, n, p, r string) error {
	if n == "" || p == "" {
		return errors.New("username and password are required")
	}
	if r == "" {
		r = "user"
	}
	hash, e := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if e != nil {
		return e
	}
	return s.repo.Create(c, model.User{Username: n, PasswordHash: string(hash), Role: r})
}
func (s *AuthServiceImpl) Login(c context.Context, n, p string) (string, error) {
	u, e := s.repo.FindByUsername(c, n)
	if e != nil {
		return "", e
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(p)) != nil {
		return "", errors.New("invalid username or password")
	}
	return u.Role, nil
}
