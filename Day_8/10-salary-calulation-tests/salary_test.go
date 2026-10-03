package salary
import "testing"

func TestCalculateSalary(t *testing.T) {

	tests := []struct {
		name       string
		basic      float64
		allowance  float64
		deduction  float64
		expected   float64
	}{
		{
			name:      "Valid salary",
			basic:     30000,
			allowance: 5000,
			deduction: 2000,
			expected:  33000,
		},
		{
			name:      "Basic salary only",
			basic:     30000,
			allowance: 0,
			deduction: 0,
			expected:  30000,
		},
		{
			name:      "No basic salary",
			basic:     0,
			allowance: 5000,
			deduction: 1000,
			expected:  4000,
		},
		{
			name:      "No allowance",
			basic:     30000,
			allowance: 0,
			deduction: 2000,
			expected:  28000,
		},
		{
			name:      "No deduction",
			basic:     30000,
			allowance: 5000,
			deduction: 0,
			expected:  35000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := CalculateSalary(
				tt.basic,
				tt.allowance,
				tt.deduction,
			)

			if result != tt.expected {
				t.Errorf(
					"Expected %.2f, but got %.2f",
					tt.expected,
					result,
				)
			}
		})
	}
}

/*
Output:
------------

PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\10-salary-calulation-tests> go test
PASS
ok      salary  1.234s
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\10-salary-calulation-tests> go test -v
=== RUN   TestCalculateSalary
=== RUN   TestCalculateSalary/Valid_salary
=== RUN   TestCalculateSalary/Basic_salary_only
=== RUN   TestCalculateSalary/No_basic_salary
=== RUN   TestCalculateSalary/No_allowance
=== RUN   TestCalculateSalary/No_deduction
--- PASS: TestCalculateSalary (0.00s)
    --- PASS: TestCalculateSalary/Valid_salary (0.00s)
    --- PASS: TestCalculateSalary/Basic_salary_only (0.00s)
    --- PASS: TestCalculateSalary/No_basic_salary (0.00s)
    --- PASS: TestCalculateSalary/No_allowance (0.00s)
    --- PASS: TestCalculateSalary/No_deduction (0.00s)
PASS
ok      salary  0.170s

*/