package utils

import "strings"

func FormatName(name string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(name)))
}
