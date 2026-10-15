package validation

import (
	"errors"
	"net/mail"
	"strings"

	"employee-api/model"
)

// These errors identify employee validation failures.
var (
	ErrNameRequired  = errors.New("employee name is required")
	ErrEmailRequired = errors.New("employee email is required")
	ErrEmailInvalid  = errors.New("employee email is invalid")
)

// ValidateEmployee checks all required employee fields.
func ValidateEmployee(employee *model.Employee) error {
	// Remove unnecessary spaces from user input.
	employee.Name = strings.TrimSpace(employee.Name)
	employee.Email = strings.TrimSpace(employee.Email)

	// Name cannot be empty.
	if employee.Name == "" {
		return ErrNameRequired
	}

	// Email cannot be empty.
	if employee.Email == "" {
		return ErrEmailRequired
	}

	// Check the email format.
	return ValidateEmail(employee.Email)
}

// ValidateEmail checks whether the email has a basic valid format.
func ValidateEmail(email string) error {
	// Reject whitespace in email addresses.
	if strings.ContainsAny(email, " \t\r\n") {
		return ErrEmailInvalid
	}

	// Parse and validate the email address syntax.
	address, err := mail.ParseAddress(email)
	if err != nil || address.Address != email {
		return ErrEmailInvalid
	}

	return nil
}