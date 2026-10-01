package assignment_task8

import "testing"

func TestFormatEmployee(t *testing.T) {
	if got := FormatEmployee("Vimal", 101); got != "Vimal #101" {
		t.Fatal(got)
	}
}

// go vet assignment_task8
// go test -v assignment_task8