package main

import "testing"

// Test data below is deliberately varied (empty name, valid name, bad
// email, valid email, etc) on purpose - a unit test's job is to check the
// function behaves correctly across different inputs, not just the one
// "happy path". This is normal test fixture data, not the business logic
// itself, which still lives in validation.go and takes real input from
// main.go at runtime.

func TestValidateName(t *testing.T) {
	if err := validateName("Anita"); err != nil {
		t.Errorf("expected no error for a valid name, got %v", err)
	}

	if err := validateName(""); err == nil {
		t.Error("expected an error for an empty name, got nil")
	}

	if err := validateName("   "); err == nil {
		t.Error("expected an error for a blank/whitespace name, got nil")
	}
}

func TestValidateEmail(t *testing.T) {
	if err := validateEmail("anita@example.com"); err != nil {
		t.Errorf("expected no error for a valid email, got %v", err)
	}

	if err := validateEmail("anita.example.com"); err == nil {
		t.Error("expected an error for an email missing @, got nil")
	}
}

func TestValidateAge(t *testing.T) {
	if err := validateAge(25); err != nil {
		t.Errorf("expected no error for age 25, got %v", err)
	}

	if err := validateAge(17); err == nil {
		t.Error("expected an error for age 17, got nil")
	}

	if err := validateAge(18); err != nil {
		t.Errorf("expected no error for age 18 (boundary case), got %v", err)
	}
}
