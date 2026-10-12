package validation

import (
	"errors"
	"strings"

	"day12/model"
)

func ValidateEmployee(e *model.Employee) error {
	if e == nil {
		return errors.New("employee data is required")
	}
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(e.Email, "@") {
		return errors.New("valid email is required")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
