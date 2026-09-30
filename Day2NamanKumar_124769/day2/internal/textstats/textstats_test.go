package textstats

import "testing"

func TestAnalyze(t *testing.T) {
	s := Analyze("Hello World 2026!")
	want := Stats{Bytes: 17, Characters: 17, NonSpace: 15, Words: 3, Vowels: 3,
		Consonants: 7, Digits: 4, Spaces: 2, Punctuation: 1, Lines: 1}
	if s != want {
		t.Fatalf("\n got %+v\nwant %+v", s, want)
	}
}

func TestUnicode(t *testing.T) {
	s := Analyze("ನಮಸ್ಕಾರ Priya")
	if s.Characters != 13 || s.Words != 2 || s.Bytes <= s.Characters {
		t.Fatalf("unicode counting wrong: %+v", s)
	}
}

func TestEmpty(t *testing.T) {
	if s := Analyze(""); s != (Stats{}) {
		t.Fatalf("expected zero stats, got %+v", s)
	}
}

func TestMultipleSpaces(t *testing.T) {
	if s := Analyze("  go   is\tfun \n"); s.Words != 3 {
		t.Fatalf("want 3 words, got %d", s.Words)
	}
}
