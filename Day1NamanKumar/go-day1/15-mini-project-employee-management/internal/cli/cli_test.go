package cli

import (
	"bytes"
	"strings"
	"testing"

	"employeeapp/internal/employee"
)

// TestEndToEnd drives the menu with scripted input.
func TestEndToEnd(t *testing.T) {
	input := strings.Join([]string{
		"1", "Asha Rao", "Engineering", "90000", // add
		"1", "Vikram", "Finance", "abc", // add with bad salary
		"3",            // display
		"2", "2", "as", // search by name
		"4", "1", // delete
		"2", "1", "1", // search deleted id
		"9", // invalid option
		"5", // exit
	}, "\n")
	var out bytes.Buffer
	New(employee.NewStore(), strings.NewReader(input), &out).Run()

	got := out.String()
	for _, want := range []string{
		"Employee added with ID 1",
		`"abc" is not a valid salary`,
		"Asha Rao",
		"Employee 1 deleted",
		"No employee found with ID 1",
		`Invalid option "9"`,
		"Goodbye!",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q", want)
		}
	}
}
