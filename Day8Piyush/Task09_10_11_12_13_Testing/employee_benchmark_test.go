// Task 13: Benchmarks
// Run: go test -bench=. -benchmem
// Run specific: go test -bench=BenchmarkValidateEmployee -benchmem

package employee

import "testing"

// Benchmark employee validation
func BenchmarkValidateEmployee(b *testing.B) {
	emp := Employee{
		Name:       "Piyush",
		Email:      "piyush@example.com",
		Age:        25,
		Department: "Engineering",
		Salary:     50000,
	}
	for i := 0; i < b.N; i++ {
		ValidateEmployee(emp)
	}
}

// Benchmark net salary calculation
func BenchmarkCalculateNetSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateNetSalary(100000, 20)
	}
}

// Benchmark annual salary
func BenchmarkCalculateAnnualSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateAnnualSalary(50000)
	}
}

// Benchmark bonus calculation
func BenchmarkCalculateBonus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateBonus(100000, 4)
	}
}

// Benchmark salary hike
func BenchmarkCalculateSalaryHike(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSalaryHike(50000, 15)
	}
}

// Benchmark tax bracket lookup
func BenchmarkGetTaxBracket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTaxBracket(800000)
	}
}

// Benchmark employee store operations
func BenchmarkEmployeeStore_Add(b *testing.B) {
	store := NewEmployeeStore()
	emp := Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}
	for i := 0; i < b.N; i++ {
		emp.Email = "p" + string(rune(i%26+97)) + "@test.com" // vary email to avoid dup
		store.Add(emp)
	}
}

// Benchmark email validation
func BenchmarkIsValidEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidEmail("piyush@example.com")
	}
}
