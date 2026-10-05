package main

import "testing"

func TestValidateEmployee(t *testing.T) {
	e := Employee{ID: 1, Name: "John", Age: 25}
	err := ValidateEmployee(e)
	if err != nil {
		t.Errorf("expected valid employee, got error: %v", err)
	}

	invalid := Employee{ID: 0, Name: "", Age: 15}
	err = ValidateEmployee(invalid)
	if err == nil {
		t.Errorf("expected error for invalid employee")
	}
}