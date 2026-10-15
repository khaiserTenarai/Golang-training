package validation

import (
	"errors"
	"net/mail"
	"strings"

	"employee-api/model"
)

var (
	ErrNameRequired  = errors.New("employee name is required")
	ErrEmailRequired = errors.New("employee email is required")
	ErrEmailInvalid  = errors.New("employee email is invalid")
)

// ValidateEmployee checks and cleans employee input.
func ValidateEmployee(employee *model.Employee) error {
	employee.Name = strings.TrimSpace(employee.Name)
	employee.Email = strings.TrimSpace(employee.Email)

	if employee.Name == "" {
		return ErrNameRequired
	}

	if employee.Email == "" {
		return ErrEmailRequired
	}

	return ValidateEmail(employee.Email)
}

// ValidateEmail checks the email format.
func ValidateEmail(email string) error {
	address, err := mail.ParseAddress(email)
	if err != nil || address.Address != email {
		return ErrEmailInvalid
	}

	return nil
}