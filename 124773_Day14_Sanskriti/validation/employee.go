package validation

import (
	"errors"
	"math"
	"net/mail"
	"strings"
	"unicode/utf8"

	"employee-api/model"
)

func ValidateEmployee(e model.Employee) error {
	name := strings.TrimSpace(e.Name)
	email := strings.TrimSpace(e.Email)
	department := strings.TrimSpace(e.Department)

	if name == "" || utf8.RuneCountInString(name) > 100 {
		return errors.New("name is required and must be at most 100 characters")
	}
	if len(email) > 150 {
		return errors.New("email must be at most 150 characters")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("valid email is required")
	}
	if department == "" || utf8.RuneCountInString(department) > 100 {
		return errors.New("department is required and must be at most 100 characters")
	}
	if e.Salary < 0 || math.IsNaN(e.Salary) || math.IsInf(e.Salary, 0) {
		return errors.New("salary must be a finite non-negative number")
	}
	return nil
}
