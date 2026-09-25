package main

import "testing"

func TestMultiply(t *testing.T) {
	// Check if 5 * 4 correctly equals 20
	if Multiply(5, 4) != 20 {
		t.Fail()
	}
}
