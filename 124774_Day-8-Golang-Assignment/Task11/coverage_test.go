package coverage

import "testing"

func TestCalculateSalary(t *testing.T) {
	got := CalculateSalary(20000, 2000)

	if got != 22000 {
		t.Errorf("got %v; want 22000", got)
	}
}
