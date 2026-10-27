package main

import "testing"

func TestAnnualSalary(t *testing.T) {
	got := AnnualSalary(20000)
	want := 240000
	if got != want {
		t.Errorf("AnnualSalary(20000) = %d; want %d", got, want)
	}
}
