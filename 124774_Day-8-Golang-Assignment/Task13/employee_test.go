package benchmark

import "testing"

func BenchmarkCalculateAnnualSalary(b *testing.B) {

	employee := Employee{
		Name:   "Muneera",
		Salary: 20000,
	}

	for i := 0; i < b.N; i++ {
		CalculateAnnualSalary(employee)
	}
}

func BenchmarkCalculateBonus(b *testing.B) {

	employee := Employee{
		Name:   "Muneera",
		Salary: 20000,
	}

	for i := 0; i < b.N; i++ {
		CalculateBonus(employee)
	}
}
