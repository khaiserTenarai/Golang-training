package main

import "testing"

func TestCalculateSalaryTableDrivenV2(t *testing.T) {
	tests := []struct {
		rate     float64
		hours    float64
		expected float64
	}{
		{20.0, 40.0, 800.0},
		{0.0, 40.0, 0.0},
		{20.0, 0.0, 0.0},
		{15.5, 10.0, 155.0},
	}

	for _, tt := range tests {
		result := CalculateSalary(tt.rate, tt.hours)
		if result != tt.expected {
			t.Errorf("CalculateSalary(%f, %f) = %f; want %f", tt.rate, tt.hours, result, tt.expected)
		}
	}
}