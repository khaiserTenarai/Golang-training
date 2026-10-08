package q7

import "fmt"

func FormatEmployee(name string, id int) string {
	if name == "" {
		name = "Unknown"
	}
	return fmt.Sprintf("%s #%d", name, id)
}

