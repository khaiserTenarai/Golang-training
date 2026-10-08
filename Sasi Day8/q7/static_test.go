package q7

import "testing"

func TestFormatEmployee(t *testing.T) {
	if got := FormatEmployee("Sasi", 101); got != "sasi #101" {
		t.Fatal(got)
	}
}

// go vet q8
// go test -v q8