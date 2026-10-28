package utility

import (
	"errors"
	"strings"

	"employee-management/model"
)

func ValidateEmployee(employee model.Employee) error {
	if strings.TrimSpace(employee.Name) == "" {
		return errors.New("employee name is required")
	}

	if !strings.Contains(employee.Email, "@") {
		return errors.New("a valid employee email is required")
	}

	if strings.TrimSpace(employee.Department) == "" {
		return errors.New("employee department is required")
	}

	if employee.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return nil
}
