package main

import (
	"testing"
)

func TestCalculateNetSalary(t *testing.T) {
	tests := []struct {
		name       string
		input      EmployeeSalary
		wantNet    float64
		wantErr    bool
	}{
		{
			name:       "valid salary with bonus and deductions",
			input:      EmployeeSalary{BaseSalary: 5000, Bonus: 1000, Deductions: 500},
			wantNet:    5500,
			wantErr:    false,
		},
		{
			name:       "valid salary with no bonus or deductions",
			input:      EmployeeSalary{BaseSalary: 4000, Bonus: 0, Deductions: 0},
			wantNet:    4000,
			wantErr:    false,
		},
		{
			name:       "negative base salary",
			input:      EmployeeSalary{BaseSalary: -3000, Bonus: 500, Deductions: 100},
			wantNet:    0,
			wantErr:    true,
		},
		{
			name:       "negative bonus",
			input:      EmployeeSalary{BaseSalary: 5000, Bonus: -200, Deductions: 100},
			wantNet:    0,
			wantErr:    true,
		},
		{
			name:       "negative deductions",
			input:      EmployeeSalary{BaseSalary: 5000, Bonus: 500, Deductions: -100},
			wantNet:    0,
			wantErr:    true,
		},
		{
			name:       "deductions greater than gross earnings",
			input:      EmployeeSalary{BaseSalary: 3000, Bonus: 500, Deductions: 4000},
			wantNet:    0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNet, err := CalculateNetSalary(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateNetSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotNet != tt.wantNet {
				t.Errorf("CalculateNetSalary() gotNet = %v, want %v", gotNet, tt.wantNet)
			}
		})
	}
}