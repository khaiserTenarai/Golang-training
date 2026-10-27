package main

import "testing"

func TestEmployeeValidate(t *testing.T) {
	employee := Employee{Name: "Ravi", Age: 22}
	if err := employee.Validate(); err != nil {
		t.Errorf("expected valid employee, got error: %v", err)
	}
}

func TestEmployeeNameRequired(t *testing.T) {
	employee := Employee{Age: 22}
	if employee.Validate() == nil {
		t.Error("expected an error when name is empty")
	}
}
