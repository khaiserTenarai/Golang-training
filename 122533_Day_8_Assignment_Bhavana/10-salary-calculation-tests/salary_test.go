package main

import "testing"

func TestNetSalaryNormalCase(t *testing.T) {
	got := netSalary(30000, 5000, 2000)
	want := 33000.0
	if got != want {
		t.Errorf("netSalary(30000, 5000, 2000) = %v, want %v", got, want)
	}
}

func TestNetSalaryNoAllowanceOrDeduction(t *testing.T) {
	got := netSalary(30000, 0, 0)
	want := 30000.0
	if got != want {
		t.Errorf("netSalary(30000, 0, 0) = %v, want %v", got, want)
	}
}

func TestNetSalaryDeductionBiggerThanAllowance(t *testing.T) {
	got := netSalary(30000, 1000, 4000)
	want := 27000.0
	if got != want {
		t.Errorf("netSalary(30000, 1000, 4000) = %v, want %v", got, want)
	}
}
