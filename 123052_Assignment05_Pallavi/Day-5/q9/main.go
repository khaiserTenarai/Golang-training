package main

import "fmt"

// Employee represents employee information
type Employee struct {
	ID   int
	Name string
}

// EmployeeRepository defines employee operations
type EmployeeRepository interface {
	AddEmployee(employee Employee)
	GetEmployee(id int) Employee
	DeleteEmployee(id int)
}

// EmployeeStore implements EmployeeRepository
type EmployeeStore struct {
	employees []Employee
}

// Add employee
func (e *EmployeeStore) AddEmployee(employee Employee) {

	e.employees = append(e.employees, employee)

	fmt.Println("Employee added successfully")
}

// Get employee
func (e *EmployeeStore) GetEmployee(id int) Employee {

	for _, employee := range e.employees {

		if employee.ID == id {
			return employee
		}
	}

	return Employee{}
}

// Delete employee
func (e *EmployeeStore) DeleteEmployee(id int) {

	for i, employee := range e.employees {

		if employee.ID == id {

			e.employees = append(
				e.employees[:i],
				e.employees[i+1:]...,
			)

			fmt.Println("Employee deleted successfully")
			return
		}
	}

	fmt.Println("Employee not found")
}

func main() {

	// Create repository
	var repository EmployeeRepository = &EmployeeStore{}

	// Add employee
	repository.AddEmployee(Employee{
		ID:   101,
		Name: "Pallavi",
	})

	// Get employee
	employee := repository.GetEmployee(101)

	fmt.Println("Employee:", employee)

	// Delete employee
	repository.DeleteEmployee(101)
}