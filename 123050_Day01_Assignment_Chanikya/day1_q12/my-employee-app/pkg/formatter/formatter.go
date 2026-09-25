// pkg/formatter/formatter.go
package formatter

import "strings"

// PrintHeader returns a capitalized title string with an underline.
func FormatHeader(title string) string {
	upperTitle := strings.ToUpper(title)
	underline := strings.Repeat("-", len(upperTitle))
	return upperTitle + "\n" + underline
}
