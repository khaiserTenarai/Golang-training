package main

import "testing"

func BenchmarkCalculateSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSalary(50000, 5000)
	}
}

//run go test -bench=.