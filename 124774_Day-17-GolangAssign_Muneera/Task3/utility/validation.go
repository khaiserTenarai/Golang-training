package utility

import (
	"errors"
	"strings"

	"employee-management/model"
)

func ValidateEmployee(e *model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)

	if e.Name == "" {
		return errors.New("employee name is required")
	}
	if e.Email == "" || !strings.Contains(e.Email, "@") {
		return errors.New("a valid employee email is required")
	}
	if e.Department == "" {
		return errors.New("employee department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
