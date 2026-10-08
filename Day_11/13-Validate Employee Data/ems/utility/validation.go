package utility

import (
	"errors"
	"strings"

	"employee-management-go/model"
)

// ValidateEmployee validates employee data.
func ValidateEmployee(employee model.Employee) error {

	if strings.TrimSpace(employee.Name) == "" {
		return errors.New("name is required")
	}

	if employee.Age < 18 || employee.Age > 60 {
		return errors.New("age must be between 18 and 60")
	}

	if !strings.Contains(employee.Email, "@") {
		return errors.New("invalid email")
	}

	return nil
}
