package utility

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func ValidateName(name string) error {
	name = CleanString(name)
	if name == "" {
		return &ValidationError{Field: "Name", Message: "name cannot be empty"}
	}
	if len(name) < 2 {
		return &ValidationError{Field: "Name", Message: "name must be at least 2 characters"}
	}
	return nil
}

func ValidateEmail(email string) error {
	email = CleanString(email)
	if email == "" {
		return &ValidationError{Field: "Email", Message: "email cannot be empty"}
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return &ValidationError{Field: "Email", Message: "email format is invalid"}
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 18 || age > 65 {
		return &ValidationError{Field: "Age", Message: "age must be between 18 and 65"}
	}
	return nil
}

func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return &ValidationError{Field: "Salary", Message: "salary must be greater than zero"}
	}
	return nil
}
