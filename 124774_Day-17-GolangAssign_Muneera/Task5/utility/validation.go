package utility

import (
	"errors"
	"strings"

	"employee-management/model"
)

func ValidateEmployee(e *model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)

	if e.Name == "" {
		return errors.New("employee name is required")
	}

	if e.Email == "" || !strings.Contains(e.Email, "@") {
		return errors.New("valid employee email is required")
	}

	if e.DepartmentID == 0 {
		return errors.New("department ID is required")
	}

	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return nil
}
