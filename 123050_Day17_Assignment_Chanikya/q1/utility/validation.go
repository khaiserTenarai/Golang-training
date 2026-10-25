package utility

import (
	"errors"
	"net/mail"
	"strings"

	"employee-management/model"
)

// ValidateEmployee checks the employee input.
func ValidateEmployee(employee *model.Employee) error {
	if employee == nil {
		return errors.New("employee cannot be nil")
	}

	employee.Name = strings.TrimSpace(employee.Name)
	employee.Email = strings.TrimSpace(employee.Email)
	employee.Department = strings.TrimSpace(employee.Department)

	if employee.Name == "" {
		return errors.New("name is required")
	}

	if _, err := mail.ParseAddress(employee.Email); err != nil {
		return errors.New("a valid email is required")
	}

	if employee.Department == "" {
		return errors.New("department is required")
	}

	if employee.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return nil
}
