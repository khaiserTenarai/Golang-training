package utility

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ":" + e.Message
}

var ErrEmployeeNotFound = errors.New("employee not found")

var ErrDuplicateEmployee = errors.New("employee already exists")

func ValidateName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return ValidationError{
			Field:   "Name",
			Message: "name cannot be empty",
		}
	}
	return nil
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return ValidationError{
			Field:   "Email",
			Message: "email cannot be empty",
		}
	}
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "Email",
			Message: "Invalid email",
		}
	}
	return nil
}

func ValidateAge(age int) error {
	if age <= 0 {
		return ValidationError{
			Field:   "Age",
			Message: "age must be greater than 0",
		}
	}
	return nil
}

func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return ValidationError{
			Field:   "Salary",
			Message: "saalry must be greater than 0",
		}
	}
	return nil
}

func ValidateEmployeee(name string, email string, age int, salary float64) error {
	err := ValidateName(name)
	if err != nil {
		return fmt.Errorf("Employee validation failed:%w", err)
	}
	err = ValidateEmail(email)
	if err != nil {
		return fmt.Errorf("Employee validation failed:%w", err)
	}

	err = ValidateAge(age)
	if err != nil {
		return fmt.Errorf("Employee validation failed:%w", err)
	}
	err = ValidateSalary(salary)
	if err != nil {
		return fmt.Errorf("Employee validation failed:%w", err)
	}
	return nil

}
