
package utility

import "strings"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func ValidateName(name string) error {
	if name == "" {
		return &ValidationError{Field: "Name", Message: "name cannot be empty"}
	}
	return nil
}
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return &ValidationError{Field: "Email", Message: "email is not valid"}
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
		return &ValidationError{Field: "Salary", Message: "salary must be greater than 0"}
	}
	return nil
}
