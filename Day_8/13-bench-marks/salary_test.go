package salary

import "testing"

// Benchmark measures the performance of CalculateSalary.
func BenchmarkCalculateSalary(b *testing.B) {

	for i := 0; i < b.N; i++ {
		CalculateSalary(30000, 5000, 2000)
	}
}

/*
Output:
------------
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy> cd .\13-bench-marks\
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\13-bench-marks> go mod init benchmak_example
go: creating new go.mod: module benchmak_example
go: to add module requirements and sums:
        go mod tidy
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\13-bench-marks> go test -bench=.
ok      benchmak_example        0.944s [no tests to run]
*/
