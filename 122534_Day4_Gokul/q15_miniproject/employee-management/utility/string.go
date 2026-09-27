package utility

import "strings"

// CleanString trims leading/trailing spaces and collapses repeated
// internal whitespace, e.g. "  Gokul   Nair " -> "Gokul Nair".
func CleanString(s string) string {
	trimmed := strings.TrimSpace(s)
	return strings.Join(strings.Fields(trimmed), " ")
}
