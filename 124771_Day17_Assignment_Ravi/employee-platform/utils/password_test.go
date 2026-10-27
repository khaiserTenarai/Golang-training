package utils

import "testing"

func TestPasswordHashAndCheck(t *testing.T) {
	hash, err := HashPassword("StrongPass123!")
	if err != nil {
		t.Fatal(err)
	}
	if !CheckPassword(hash, "StrongPass123!") {
		t.Fatal("expected password to match")
	}
	if CheckPassword(hash, "wrong") {
		t.Fatal("wrong password should not match")
	}
}
