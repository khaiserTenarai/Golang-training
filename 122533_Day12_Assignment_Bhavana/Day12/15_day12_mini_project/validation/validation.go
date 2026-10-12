package validation

import (
	"net/mail"
	"strings"

	"employeeapi/model"
)

// ValidateEmployee checks every rule and returns all the problems it
// finds, or an empty list if the employee is fine.
func ValidateEmployee(e model.Employee) []string {
	var problems []string

	if len(strings.TrimSpace(e.Name)) < 2 {
		problems = append(problems, "name must have at least 2 characters")
	}
	if addr, err := mail.ParseAddress(e.Email); err != nil || addr.Address != e.Email {
		problems = append(problems, "email is not a valid address")
	}
	if strings.TrimSpace(e.Department) == "" {
		problems = append(problems, "department is required")
	}
	if e.Salary <= 0 {
		problems = append(problems, "salary must be greater than 0")
	}
	return problems
}
