package strutil

import "testing"

// Benchmark string concatenation
func BenchmarkConcatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatString(1000)
	}
}

// Benchmark strings.Builder
func BenchmarkBuilderString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderString(1000)
	}
}