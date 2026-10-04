package main

import "testing"

func TestAddEmployee(t *testing.T) {
	service := EmployeeService{}

	employee := Employee{
		ID:     1,
		Name:   "Sanskriti",
		Age:    22,
		Salary: 50000,
	}

	err := service.AddEmployee(employee)

	if err != nil {
		t.Errorf("Expected employee to be added, got error: %v", err)
	}
}

func TestAddEmployeeInvalidAge(t *testing.T) {
	service := EmployeeService{}

	employee := Employee{
		ID:     1,
		Name:   "Sanskriti",
		Age:    0,
		Salary: 50000,
	}

	err := service.AddEmployee(employee)

	if err == nil {
		t.Error("Expected error for invalid age")
	}
}

func TestGetEmployee(t *testing.T) {
	service := EmployeeService{}

	employee := Employee{
		ID:     1,
		Name:   "Sanskriti",
		Age:    22,
		Salary: 50000,
	}

	service.AddEmployee(employee)

	result, err := service.GetEmployee(1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result.Name != "Sanskriti" {
		t.Errorf("Expected Sanskriti, got %s", result.Name)
	}
}

func TestGetEmployeeNotFound(t *testing.T) {
	service := EmployeeService{}

	_, err := service.GetEmployee(100)

	if err == nil {
		t.Error("Expected employee not found error")
	}
}

func TestRace(t *testing.T) {
	service := EmployeeService{}

	done := make(chan bool)

	go func() {
		employee := Employee{
			ID:     1,
			Name:   "Sanskriti",
			Age:    22,
			Salary: 50000,
		}

		service.AddEmployee(employee)

		done <- true
	}()

	go func() {
		service.GetAllEmployees()

		done <- true
	}()

	<-done
	<-done
}
