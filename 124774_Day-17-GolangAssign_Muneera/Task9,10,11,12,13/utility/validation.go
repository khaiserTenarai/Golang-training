package utility

import (
	"fmt"
	"net/mail"
	"strings"

	"employee-management/model"
)

func ValidateEmployee(e model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)

	if e.Name == "" {
		return fmt.Errorf("employee name is required")
	}

	if e.Email == "" {
		return fmt.Errorf("employee email is required")
	}

	if _, err := mail.ParseAddress(e.Email); err != nil {
		return fmt.Errorf("invalid employee email")
	}

	if e.Department == "" {
		return fmt.Errorf("employee department is required")
	}

	if e.Salary < 0 {
		return fmt.Errorf("salary cannot be negative")
	}

	return nil
}
