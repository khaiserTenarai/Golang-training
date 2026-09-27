// Package utility has small, reusable helper functions - validation
// rules and string cleanup - used by the service layer.
package utility

import (
	"fmt"
	"strings"
)

// ValidationError is a custom error type so callers can tell exactly
// which field failed, instead of just getting a plain string back.
type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

// ValidateName checks that a name isn't empty or just whitespace.
func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return &ValidationError{Field: "Name", Message: "cannot be empty"}
	}
	return nil
}

// ValidateEmail does a very basic sanity check - good enough for this
// assignment, not meant to be a full RFC-grade email validator.
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return &ValidationError{Field: "Email", Message: "must be a valid email address"}
	}
	return nil
}

// ValidateAge makes sure the employee is a working adult and not some
// unrealistic age.
func ValidateAge(age int) error {
	if age < 18 || age > 65 {
		return &ValidationError{Field: "Age", Message: "must be between 18 and 65"}
	}
	return nil
}

// ValidateSalary rejects negative or zero salary values.
func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return &ValidationError{Field: "Salary", Message: "must be greater than zero"}
	}
	return nil
}
