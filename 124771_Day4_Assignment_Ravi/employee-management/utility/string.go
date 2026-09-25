package utility

import "strings"

// Removes leading and trailing whitespace.
func CleanString(value string) string {
	return strings.TrimSpace(value)
}
