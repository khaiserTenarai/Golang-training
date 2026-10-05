package main

import "testing"

// Only 3 of the 4 Grade() branches are tested here on purpose (A, B, and
// the default F case) - see README.md for what that means for coverage,
// and the C case is left as an exercise to add.

func TestGradeA(t *testing.T) {
	if got := Grade(95); got != "A" {
		t.Errorf("Grade(95) = %s, want A", got)
	}
}

func TestGradeB(t *testing.T) {
	if got := Grade(80); got != "B" {
		t.Errorf("Grade(80) = %s, want B", got)
	}
}

func TestGradeF(t *testing.T) {
	if got := Grade(30); got != "F" {
		t.Errorf("Grade(30) = %s, want F", got)
	}
}
