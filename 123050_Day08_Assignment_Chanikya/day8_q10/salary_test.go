package main

import "testing"

func TestCalculateSalary(t *testing.T) {

	basicSalary := 50000.0
	bonus := 5000.0

	expected := 55000.0

	result := CalculateSalary(basicSalary, bonus)

	if result != expected {
		t.Errorf(
			"expected %.2f, got %.2f",
			expected,
			result,
		)
	}
}
