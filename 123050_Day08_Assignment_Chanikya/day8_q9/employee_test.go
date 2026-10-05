package main

import "testing"

func TestValidateEmployee(t *testing.T) {

	employee := Employee{
		Name:   "ram",
		Salary: 50000,
	}

	err := ValidateEmployee(employee)

	if err != nil {
		t.Errorf("expected employee to be valid, got: %v", err)
	}
}
func TestValidateEmployeeWithoutName(t *testing.T) {

	employee := Employee{
		Name:   "",
		Salary: 50000,
	}

	err := ValidateEmployee(employee)

	if err == nil {
		t.Error("expected error for empty employee name")
	}
}
func TestValidateEmployeeWithInvalidSalary(t *testing.T) {

	employee := Employee{
		Name:   "ram",
		Salary: 0,
	}

	err := ValidateEmployee(employee)

	if err == nil {
		t.Error("expected error for invalid salary")
	}
}
