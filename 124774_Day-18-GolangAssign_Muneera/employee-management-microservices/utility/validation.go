package utility

import (
	"fmt"
	"math"
	"net/mail"
	"strings"
)

// ValidateEmployee checks employee details.
func ValidateEmployee(
	name, email string,
	departmentID uint,
	salary float64,
) error {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || len(name) > 100 {
		return fmt.Errorf("name is required and must be at most 100 characters")
	}

	if len(email) == 0 || len(email) > 150 {
		return fmt.Errorf("email is required and must be at most 150 characters")
	}

	address, err := mail.ParseAddress(email)
	if err != nil || address.Address != email {
		return fmt.Errorf("invalid email address")
	}

	if departmentID == 0 {
		return fmt.Errorf("department ID must be greater than zero")
	}

	if math.IsNaN(salary) || math.IsInf(salary, 0) || salary < 0 {
		return fmt.Errorf("salary must be a valid non-negative number")
	}

	return nil
}

// ValidateDepartment checks department details.
func ValidateDepartment(name, description string) error {
	name = strings.TrimSpace(name)

	if name == "" || len(name) > 100 {
		return fmt.Errorf("department name is required and must be at most 100 characters")
	}

	if len(description) > 500 {
		return fmt.Errorf("description must be at most 500 characters")
	}

	return nil
}
