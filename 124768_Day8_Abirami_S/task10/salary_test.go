package main

import "testing"

func TestCalculateSalary(t *testing.T) {
	result := calculateSalary(70000)
	expected := 77000
	if result != float64(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
