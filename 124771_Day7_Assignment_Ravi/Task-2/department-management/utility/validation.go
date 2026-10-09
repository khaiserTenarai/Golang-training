package utility

import "strings"

func ValidateDepartmentName(name string) bool {
	name = strings.TrimSpace(name)

	return len(name) >= 2
}

func ValidateEmployeeName(name string) bool {
	name = strings.TrimSpace(name)

	return len(name) >= 2
}

func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)

	return strings.Contains(email, "@") &&
		strings.Contains(email, ".")
}

func ValidateSalary(salary float64) bool {
	return salary >= 0
}

func ValidateID(id int) bool {
	return id > 0
}
