package salary
import "testing"

func TestCalculateSalary(t *testing.T) {

	tests := []struct {
		name      string
		basic     float64
		allowance float64
		deduction float64
		expected  float64
	}{
		{"Valid salary", 30000, 5000, 2000, 33000},
		{"Basic salary only", 30000, 0, 0, 30000},
		{"No allowance", 30000, 0, 2000, 28000},
		{"No deduction", 30000, 5000, 0, 35000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := CalculateSalary(
				tt.basic,
				tt.allowance,
				tt.deduction,
			)

			if result != tt.expected {
				t.Errorf("Expected %.2f, but got %.2f",
					tt.expected, result)
			}
		})
	}
}

/*
Output :
---------
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go mod init test-coverage
go: creating new go.mod: module test-coverage
go: to add module requirements and sums:
        go mod tidy
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go test
PASS
ok      test-coverage   1.064s
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go test -v
=== RUN   TestCalculateSalary
=== RUN   TestCalculateSalary/Valid_salary
=== RUN   TestCalculateSalary/Basic_salary_only
=== RUN   TestCalculateSalary/No_allowance
=== RUN   TestCalculateSalary/No_deduction
--- PASS: TestCalculateSalary (0.00s)
    --- PASS: TestCalculateSalary/Valid_salary (0.00s)
    --- PASS: TestCalculateSalary/Basic_salary_only (0.00s)
    --- PASS: TestCalculateSalary/No_allowance (0.00s)
    --- PASS: TestCalculateSalary/No_deduction (0.00s)
PASS
ok      test-coverage   0.164s


PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go version
go version go1.27.1 windows/amd64
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go test -coverprofile="coverage.out" .
ok      test-coverage   1.227s  coverage: 100.0% of statements
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go tool cover -func="coverage.out"
test-coverage/salary.go:2:      CalculateSalary 100.0%
total:                          (statements)    100.0%
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\11-measure-test-coverage> go tool cover -html="coverage.out"


all cmds sequence:
-------------
go test
go test -v
go test -cover
go test -coverprofile="coverage.out" .
go tool cover -func="coverage.out"
go tool cover -html="coverage.out"
*/