package main

import "testing"

func TestCalculateSalary(t *testing.T) {
	result := CalculateSalary(50000, 5000)

	expected := 55000.0

	if result != expected {
		t.Errorf(
			"Expected %.2f but got %.2f",
			expected,
			result,
		)
	}
}