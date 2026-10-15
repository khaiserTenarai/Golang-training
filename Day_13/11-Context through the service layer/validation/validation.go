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

// ValidateEmployee validates and normalizes employee input.
func ValidateEmployee(employee *model.Employee) error {
	employee.Name = strings.TrimSpace(employee.Name)
	employee.Email = strings.TrimSpace(employee.Email)

	if employee.Name == "" {
		return ErrNameRequired
	}

	if employee.Email == "" {
		return ErrEmailRequired
	}

	if err := ValidateEmail(employee.Email); err != nil {
		return err
	}

	return nil
}

// ValidateEmail checks whether the email address has a valid basic format.
func ValidateEmail(email string) error {
	if strings.ContainsAny(email, " \t\r\n") {
		return ErrEmailInvalid
	}

	address, err := mail.ParseAddress(email)
	if err != nil || address.Address != email {
		return ErrEmailInvalid
	}

	return nil
}