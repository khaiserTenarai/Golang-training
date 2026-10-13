package validation

import (
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"

	"employee-api-go/internal/domain"
)

func ValidateCreateEmployee(input domain.CreateEmployeeInput) error {
	return validateEmployee(input.Name, input.Email, input.Department)
}

func ValidateUpdateEmployee(input domain.UpdateEmployeeInput) error {
	return validateEmployee(input.Name, input.Email, input.Department)
}

func validateEmployee(name, email, department string) error {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	department = strings.TrimSpace(department)

	if name == "" || utf8.RuneCountInString(name) > 100 {
		return errors.New("name is required and must be at most 100 characters")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("email must be a valid email address")
	}
	if len(email) > 254 {
		return errors.New("email must be at most 254 characters")
	}
	if department == "" || utf8.RuneCountInString(department) > 80 {
		return errors.New("department is required and must be at most 80 characters")
	}
	return nil
}
