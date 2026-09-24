package main

import "testing"

func TestGreet(t *testing.T) {
	if got := Greet("Go"); got != "Hello, Go!" {
		t.Errorf("Greet = %q", got)
	}
}
