package main

import "testing"

func Square(n int) int {
	return n * n
}

func TestSquare(t *testing.T) {
	if Square(6) != 36 {
		t.Fail()
	}
}