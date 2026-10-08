package q13

import "testing"

func TestContains(t *testing.T) {
	if !Contains("hello golang", "golang") {
		t.Fatal("not found")
	}
}

func BenchmarkContains(b *testing.B) {
	text := "Go is a fast, simple, productive programming language."
	for i := 0; i < b.N; i++ {
		Contains(text, "productive")
	}
}

func BenchmarkBuildMessage(b *testing.B) {
	words := []string{"Go", "makes", "concurrent", "programming", "simple"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildMessage(words)
	}
}