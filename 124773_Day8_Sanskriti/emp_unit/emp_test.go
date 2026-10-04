package main

import "testing"

func TestValidateName(t *testing.T) {
	err := ValidateName("Sanskriti")

	if err != nil {
		t.Error("Name should be valid")
	}
}

func TestValidateAge(t *testing.T) {
	err := ValidateAge(22)

	if err != nil {
		t.Error("Age should be valid")
	}
}

func TestValidateSalary(t *testing.T) {
	err := ValidateSalary(50000)

	if err != nil {
		t.Error("Salary should be valid")
	}
}