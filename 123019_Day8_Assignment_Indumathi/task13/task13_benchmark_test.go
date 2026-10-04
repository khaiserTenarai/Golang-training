package main

import "testing"

func BenchmarkCalculateSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSalary(25.0, 160.0)
	}
}