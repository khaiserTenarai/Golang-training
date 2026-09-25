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
	if strings.TrimSpace(name) == "" {
		return &ValidationError{
			Field:   "Name",
			Message: "name cannot be empty",
		}
	}

	return nil
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {
		return &ValidationError{
			Field:   "Email",
			Message: "invalid email",
		}
	}

	return nil
}

func ValidateAge(age int) error {
	if age < 18 || age > 100 {
		return &ValidationError{
			Field:   "Age",
			Message: "age must be between 18 and 100",
		}
	}

	return nil
}

func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return &ValidationError{
			Field:   "Salary",
			Message: "salary must be greater than zero",
		}
	}

	return nil
}
