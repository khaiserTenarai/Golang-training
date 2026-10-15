package validation

import (
	"errors"
	"net/mail"
	"strings"

	"employee-api/model"
)

// These errors identify validation failures.
var (
	ErrNameRequired  = errors.New("employee name is required")
	ErrEmailRequired = errors.New("employee email is required")
	ErrEmailInvalid  = errors.New("employee email is invalid")
)

// ValidateEmployee validates and cleans employee input.
func ValidateEmployee(employee *model.Employee) error {
	// Remove unnecessary spaces.
	employee.Name = strings.TrimSpace(employee.Name)
	employee.Email = strings.TrimSpace(employee.Email)

	// Name is required.
	if employee.Name == "" {
		return ErrNameRequired
	}

	// Email is required.
	if employee.Email == "" {
		return ErrEmailRequired
	}

	// Validate the email format.
	return ValidateEmail(employee.Email)
}

// ValidateEmail checks the basic email address format.
func ValidateEmail(email string) error {
	// Reject whitespace in email addresses.
	if strings.ContainsAny(email, " \t\r\n") {
		return ErrEmailInvalid
	}

	// Parse the address and reject display-name formats.
	address, err := mail.ParseAddress(email)
	if err != nil || address.Address != email {
		return ErrEmailInvalid
	}

	return nil
}