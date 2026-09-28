package utility

import (
	"errors"
	"strings"
	"unicode"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("Name can't be empty")
	}
	for _, ch := range name {
		if !unicode.IsLetter(ch) && ch != ' ' {
			return errors.New("Name should contain only letters")
		}
	}
	return nil
}
func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("Email can't be empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("Invalid email")
	}
	return nil
}
func ValidateAge(age int) error {
	if age < 18 {
		return errors.New("Age must be 18 or above")
	}
	return nil
}
func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return errors.New("Salary must be greater than 0")
	}
	return nil
}
