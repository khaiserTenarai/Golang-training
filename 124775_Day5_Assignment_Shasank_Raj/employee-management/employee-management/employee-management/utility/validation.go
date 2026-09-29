package utility

import "strings"

func IsValidName(name string) bool {
	return strings.TrimSpace(name) != ""
}

func IsValidAge(age int) bool {
	return age > 0 && age <= 100
}

func IsValidSalary(salary float64) bool {
	return salary >= 0
}

func IsValidEmail(email string) bool {
	return strings.Contains(email, "@")
}