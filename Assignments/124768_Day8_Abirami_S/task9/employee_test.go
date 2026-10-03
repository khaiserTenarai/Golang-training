package main

import "testing"

func TestValidateEmployee(t *testing.T) {
	employee := Employee{
		ID:     101,
		Name:   "Tom",
		Salary: 70000,
	}
	result := validateEmployee(employee)
	if !result {
		t.Error("Invalid employee")
	}
}
