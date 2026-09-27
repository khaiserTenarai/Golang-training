package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

// EmployeeRepository interface
type EmployeeRepository interface {
	AddEmployee()
	SearchEmployee()
	DeleteEmployee()
}

// Implementing the interface
type EmployeeRepo struct {
}

// AddEmployee method
func (e EmployeeRepo) AddEmployee() {
	fmt.Println("Employee added successfully")
}

// SearchEmployee method
func (e EmployeeRepo) SearchEmployee() {
	fmt.Println("Employee searched successfully")
}

// DeleteEmployee method
func (e EmployeeRepo) DeleteEmployee() {
	fmt.Println("Employee deleted successfully")
}

func main() {

	var repo EmployeeRepository = EmployeeRepo{}

	repo.AddEmployee()
	repo.SearchEmployee()
	repo.DeleteEmployee()
}
