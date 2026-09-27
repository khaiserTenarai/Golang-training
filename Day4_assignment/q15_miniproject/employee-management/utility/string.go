package utility

import "strings"

func CleanString(s string) string {
	trimmed := strings.TrimSpace(s)
	return strings.Join(strings.Fields(trimmed), " ")
}
