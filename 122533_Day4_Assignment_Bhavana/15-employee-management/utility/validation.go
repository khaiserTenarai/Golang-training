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
	if name == "" {
		return &ValidationError{Field: "Name", Message: "name cannot be empty"}
	}
	return nil
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "Email", Message: "email must contain @"}
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "Age", Message: "age must be 18 or above"}
	}
	return nil
}

func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return &ValidationError{Field: "Salary", Message: "salary must be greater than 0"}
	}
	return nil
}
