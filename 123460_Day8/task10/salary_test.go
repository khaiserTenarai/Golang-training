package employee

import (
	"testing"
)

func TestCalculateNetSalary(t *testing.T) {
	tests := []struct {
		name        string
		payroll     Payroll
		wantSalary  float64
		wantErr     bool
	}{
		{
			name: "Standard Salary with Bonus and Tax",
			payroll: Payroll{
				BaseSalary: 5000.0,
				Bonus:      1000.0,
				Deductions: 500.0,
				TaxRate:    0.20, // (6000 * 0.20 = 1200 tax). 6000 - 1200 - 500 = 4300
			},
			wantSalary: 4300.0,
			wantErr:    false,
		},
		{
			name: "Zero Bonus and Zero Deductions",
			payroll: Payroll{
				BaseSalary: 4000.0,
				Bonus:      0.0,
				Deductions: 0.0,
				TaxRate:    0.10, // (4000 * 0.10 = 400 tax). 4000 - 400 = 3600
			},
			wantSalary: 3600.0,
			wantErr:    false,
		},
		{
			name: "Negative Base Salary",
			payroll: Payroll{
				BaseSalary: -1000.0,
				Bonus:      0.0,
				Deductions: 0.0,
				TaxRate:    0.20,
			},
			wantSalary: 0.0,
			wantErr:    true,
		},
		{
			name: "Invalid Tax Rate (> 1)",
			payroll: Payroll{
				BaseSalary: 5000.0,
				Bonus:      0.0,
				Deductions: 0.0,
				TaxRate:    1.2,
			},
			wantSalary: 0.0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.payroll.CalculateNetSalary()
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateNetSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.wantSalary {
				t.Errorf("CalculateNetSalary() = %v, want %v", got, tt.wantSalary)
			}
		})
	}
}