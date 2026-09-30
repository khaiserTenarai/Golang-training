package employee

import (
	"errors"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Email      string
	Department string
	Age        int
	BaseSalary float64
}

var (
	ErrInvalidID     = errors.New("id must be positive")
	ErrEmptyName     = errors.New("name is required")
	ErrInvalidEmail  = errors.New("email is invalid")
	ErrInvalidAge    = errors.New("age must be between 18 and 65")
	ErrInvalidSalary = errors.New("salary must be positive")
)

// Validate checks all employee fields and returns the first problem found.
func (e Employee) Validate() error {
	switch {
	case e.ID <= 0:
		return ErrInvalidID
	case strings.TrimSpace(e.Name) == "":
		return ErrEmptyName
	case !strings.Contains(e.Email, "@") || strings.HasPrefix(e.Email, "@") || strings.HasSuffix(e.Email, "@"):
		return ErrInvalidEmail
	case e.Age < 18 || e.Age > 65:
		return ErrInvalidAge
	case e.BaseSalary <= 0:
		return ErrInvalidSalary
	}
	return nil
}