package main

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1000)
	}
}

//go test -bench=. math.go math_test.go