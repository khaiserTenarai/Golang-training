package main

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"
)

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EmployeeInput struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type ListParams struct {
	Page       int
	Limit      int
	Department string
	Search     string
	Sort       string
	Order      string
}

type Repository interface {
	Create(ctx context.Context, in EmployeeInput) (Employee, error)
	GetByID(ctx context.Context, id int) (Employee, error)
	List(ctx context.Context, p ListParams) ([]Employee, int, error)
	Update(ctx context.Context, id int, in EmployeeInput) (Employee, error)
	Delete(ctx context.Context, id int) error
	Ping(ctx context.Context) error
}

var (
	ErrNotFound       = errors.New("employee not found")
	ErrDuplicateEmail = errors.New("email already exists")
)

var allowedDepartments = map[string]bool{
	"Engineering": true,
	"Finance":     true,
	"HR":          true,
	"Marketing":   true,
	"Sales":       true,
}

func validateInput(in *EmployeeInput) map[string]string {
	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	in.Department = strings.TrimSpace(in.Department)

	fields := map[string]string{}

	if n := utf8.RuneCountInString(in.Name); n < 2 || n > 100 {
		fields["name"] = "must be between 2 and 100 characters"
	}

	addr, err := mail.ParseAddress(in.Email)
	if err != nil || addr.Address != in.Email {
		fields["email"] = "must be a valid email address"
	}

	if !allowedDepartments[in.Department] {
		fields["department"] = "must be one of: Engineering, Finance, HR, Marketing, Sales"
	}

	if in.Salary <= 0 || in.Salary > 10000000 {
		fields["salary"] = "must be greater than 0 and at most 10000000"
	}

	return fields
}
