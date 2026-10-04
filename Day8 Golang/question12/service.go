package main

import (
	"errors"
	"strings"
)

type UserService interface {
	RegisterUser(name, email string) (*User, error)
	FetchActiveUsers(page, pageSize int) (*PaginatedResponse, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(name, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if email == "" || !strings.Contains(email, "@") {
		return nil, errors.New("invalid email address")
	}

	user := &User{
		Name:     name,
		Email:    email,
		IsActive: true,
	}

	err := s.repo.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) FetchActiveUsers(page, pageSize int) (*PaginatedResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	users, totalCount, err := s.repo.GetActiveUsers(pageSize, offset)
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []User{}
	}

	return &PaginatedResponse{
		Data:       users,
		TotalCount: totalCount,
		Page:       page,
	}, nil
}