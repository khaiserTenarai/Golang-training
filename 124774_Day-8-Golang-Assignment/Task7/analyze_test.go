package static_analysis

import "testing"

func TestFormatEmployee(t *testing.T) {
	if got := FormatEmployee("Muneera", 101); got != "Muneera #101" {
		t.Fatal(got)
	}
}
