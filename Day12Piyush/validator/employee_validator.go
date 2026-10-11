// Task 9: Validation Layer — separated from handler and service.
// Keeps validation logic reusable and testable.

package validator

import (
	"strings"

	"production_api/models"
)

func ValidateEmployee(emp models.Employee) []models.ValidationError {
	var errs []models.ValidationError

	if strings.TrimSpace(emp.Name) == "" {
		errs = append(errs, models.ValidationError{Field: "name", Message: "Name is required"})
	} else if len(emp.Name) < 2 {
		errs = append(errs, models.ValidationError{Field: "name", Message: "Name must be at least 2 characters"})
	} else if len(emp.Name) > 100 {
		errs = append(errs, models.ValidationError{Field: "name", Message: "Name must not exceed 100 characters"})
	}

	if strings.TrimSpace(emp.Email) == "" {
		errs = append(errs, models.ValidationError{Field: "email", Message: "Email is required"})
	} else if !strings.Contains(emp.Email, "@") || !strings.Contains(emp.Email, ".") {
		errs = append(errs, models.ValidationError{Field: "email", Message: "Invalid email format"})
	}

	if emp.Age < 18 || emp.Age > 65 {
		errs = append(errs, models.ValidationError{Field: "age", Message: "Age must be between 18 and 65"})
	}

	if strings.TrimSpace(emp.Department) == "" {
		errs = append(errs, models.ValidationError{Field: "department", Message: "Department is required"})
	}

	if emp.Salary < 0 {
		errs = append(errs, models.ValidationError{Field: "salary", Message: "Salary cannot be negative"})
	}

	validRoles := map[string]bool{"developer": true, "designer": true, "marketer": true, "lead": true, "hr": true, "admin": true, "intern": true}
	if emp.Role != "" && !validRoles[strings.ToLower(emp.Role)] {
		errs = append(errs, models.ValidationError{Field: "role", Message: "Invalid role"})
	}

	return errs
}
