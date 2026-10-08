package q13

import "strings"

func Contains(text, target string) bool {
	return strings.Contains(text, target)
}

func BuildMessage(words []string) string {
	var b strings.Builder
	for _, w := range words {
		b.WriteString(w)
		b.WriteByte(' ')
	}
	return strings.TrimSpace(b.String())
}