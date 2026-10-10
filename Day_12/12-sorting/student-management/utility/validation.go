package utility

import "strings"

func ValidStudent(name, grade string) bool {
	return strings.TrimSpace(name) != "" && strings.TrimSpace(grade) != ""
}
