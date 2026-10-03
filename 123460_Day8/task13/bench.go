package strutil

import "strings"

// ConcatString uses standard += operator for string concatenation
func ConcatString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

// BuilderString uses strings.Builder for efficient concatenation
func BuilderString(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString("a")
	}
	return builder.String()
}