package main

import "testing"

func TestCalculateSalary(t *testing.T) {

	tests := []struct {
		name        string
		basicSalary float64
		bonus       float64
		expected    float64
	}{
		{
			name:        "Normal salary",
			basicSalary: 50000,
			bonus:       5000,
			expected:    55000,
		},
		{
			name:        "Higher salary",
			basicSalary: 80000,
			bonus:       10000,
			expected:    90000,
		},
		{
			name:        "No bonus",
			basicSalary: 40000,
			bonus:       0,
			expected:    40000,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result := CalculateSalary(
				test.basicSalary,
				test.bonus,
			)

			if result != test.expected {
				t.Errorf(
					"expected %.2f, got %.2f",
					test.expected,
					result,
				)
			}
		})
	}
}
