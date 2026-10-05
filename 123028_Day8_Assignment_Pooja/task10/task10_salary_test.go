package main

import "testing"

func TestCalculateSalary(t *testing.T) {
	salary := CalculateSalary(20.0, 40.0)
	expected := 800.0

	if salary != expected {
		t.Errorf("expected %f, got %f", expected, salary)
	}
}