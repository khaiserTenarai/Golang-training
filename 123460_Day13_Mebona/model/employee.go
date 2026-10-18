package model

import (
	"net/mail"
	"strings"
	"time"
)

const MaxSalary = 10_000_000

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"createdAt"`
}

// EmployeeInput is the body of POST and PUT requests.
type EmployeeInput struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// Validate cleans the input (trims spaces, lower-cases the email) and checks the rules.
func (in *EmployeeInput) Validate() error {
	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	in.Department = strings.TrimSpace(in.Department)

	fields := map[string]string{}
	if n := len([]rune(in.Name)); n < 2 || n > 100 {
		fields["name"] = "must be 2 to 100 characters"
	}
	if addr, err := mail.ParseAddress(in.Email); err != nil || addr.Address != in.Email {
		fields["email"] = "must be a valid email address"
	}
	if n := len([]rune(in.Department)); n < 1 || n > 50 {
		fields["department"] = "must be 1 to 50 characters"
	}
	if in.Salary < 0 || in.Salary > MaxSalary {
		fields["salary"] = "must be between 0 and 10000000"
	}
	if len(fields) > 0 {
		return &ValidationError{Fields: fields}
	}
	return nil
}

// ListFilter is what the repository understands. Limit 0 means "no limit".
type ListFilter struct {
	Search string
	Limit  int
	Offset int
}
