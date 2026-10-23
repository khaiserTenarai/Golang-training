
package utility

import (
	"fmt"
	"net/mail"
	"strings"

	"secure-gin-api/model"
)

// ValidateEmployee checks employee input before saving it.
func ValidateEmployee(e model.EmployeeInput) error {
	// Validate name.
	name := strings.TrimSpace(e.Name)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if len(name) > 100 {
		return fmt.Errorf("name must be at most 100 characters")
	}

	// Validate email.
	email := strings.TrimSpace(e.Email)
	if email == "" {
		return fmt.Errorf("email is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("valid email is required")
	}

	// Validate department.
	department := strings.TrimSpace(e.Department)
	if department == "" {
		return fmt.Errorf("department is required")
	}
	if len(department) > 100 {
		return fmt.Errorf("department must be at most 100 characters")
	}

	// Validate salary.
	if e.Salary <= 0 {
		return fmt.Errorf("salary must be greater than zero")
	}

	return nil
}
