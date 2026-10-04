package assignment_task8

import "testing"

func TestSafeCounter(t *testing.T) {
	if got := SafeCounter(10); got != 10000 {
		t.Fatalf("got %d; want 10000", got)
	}
}