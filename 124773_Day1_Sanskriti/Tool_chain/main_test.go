package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 5
	actual := Add(2, 3) // Calls Add() from main.go directly
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
