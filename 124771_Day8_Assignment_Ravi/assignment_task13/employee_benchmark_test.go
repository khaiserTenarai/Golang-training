package main

import "testing"

func addNumbers(a, b int) int {
	return a + b
}

func BenchmarkAddNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addNumbers(10, 20)
	}
}
