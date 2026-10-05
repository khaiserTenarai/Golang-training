package main

import "testing"

// a normal test - just checks the answer is correct
func TestFactorial(t *testing.T) {
	if got := factorial(5); got != 120 {
		t.Errorf("factorial(5) = %d, want 120", got)
	}
}

// a benchmark - measures how fast factorial(10) runs. Go automatically
// calls this function b.N times and reports the average time per call.
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(10)
	}
}
