package employee

import "testing"

func BenchmarkNetSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NetSalary(50000, 10, 20)
	}
}

func BenchmarkValidate(b *testing.B) {
	e := Employee{ID: 1, Name: "Asha", Email: "a@x.com", Age: 30, BaseSalary: 50000}
	for i := 0; i < b.N; i++ {
		_ = e.Validate()
	}
}