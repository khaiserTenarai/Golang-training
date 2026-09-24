// Package utils contains small, reusable helper functions
// that are not tied to any business domain.
package utils

import (
	"fmt"
	"strings"
)

// TitleCase converts "asha rao" to "Asha Rao".
func TitleCase(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}

// FormatCurrency formats an amount in Indian Rupees, e.g. "₹12,50,000.00".
func FormatCurrency(amount float64) string {
	whole := fmt.Sprintf("%.2f", amount)
	intPart, frac := whole[:len(whole)-3], whole[len(whole)-3:]
	// Indian grouping: last 3 digits, then groups of 2.
	if len(intPart) > 3 {
		head, tail := intPart[:len(intPart)-3], intPart[len(intPart)-3:]
		var groups []string
		for len(head) > 2 {
			groups = append([]string{head[len(head)-2:]}, groups...)
			head = head[:len(head)-2]
		}
		if head != "" {
			groups = append([]string{head}, groups...)
		}
		intPart = strings.Join(groups, ",") + "," + tail
	}
	return "₹" + intPart + frac
}
