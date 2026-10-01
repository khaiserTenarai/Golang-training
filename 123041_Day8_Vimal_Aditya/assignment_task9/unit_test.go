package assignment_task9

import "testing"

func TestValidateEmployee(t *testing.T) {
	emp := Employee{ID: 101, Name: "Vimal", Salary: 500000}
	err := ValidateEmployee(emp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateEmployeeInvalidID(t *testing.T) {
	emp := Employee{ID: 0, Name: "Vimal", Salary: 500000}
	_, err := ValidateEmployee(emp), ValidateEmployee(emp)
	if err == nil {
		t.Fatal("expected error")
	}
}