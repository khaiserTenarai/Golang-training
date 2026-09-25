package utility

import "strings"

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func ValidateName(name string) error {
	if name == "" {
		return ValidationError{
			Field:   "Name",
			Message: "name cannot be empty",
		}
	}

	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return ValidationError{
			Field:   "Email",
			Message: "email can not be empty",
		}
	}
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "Email",
			Message: "invalid email format",
		}
	}

	return nil
}

func ValidateAge(age int) error {
	if age < 18 {
		return ValidationError{
			Field:   "Age",
			Message: "age must be least 18",
		}
	}
	return nil
}

func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return ValidationError{
			Field:   "Salary",
			Message: "Salary must be more than zero",
		}
	}
	return nil
}
