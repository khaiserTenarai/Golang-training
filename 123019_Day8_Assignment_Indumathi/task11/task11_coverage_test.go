package main

import "testing"

func TestCoverage(t *testing.T) {
	if CalculateSalary(0, 40) != 0 {
		t.Error("expected 0")
	}
	if CalculateSalary(20, 0) != 0 {
		t.Error("expected 0")
	}
	if CalculateSalary(20, 40) != 800 {
		t.Error("expected 800")
	}
}