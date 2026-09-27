package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

// Interface
type EmployeeRepository interface {
	Add(employee Employee)
	Get(id int) Employee
}

// Implementation
type EmployeeRepositoryImpl struct {
	employees []Employee
}

func (r *EmployeeRepositoryImpl) Add(employee Employee) {
	r.employees = append(r.employees, employee)
}

func (r *EmployeeRepositoryImpl) Get(id int) Employee {
	for _, employee := range r.employees {
		if employee.ID == id {
			return employee
		}
	}

	return Employee{}
}

func main() {
	repository := EmployeeRepositoryImpl{}

	repository.Add(Employee{
		ID:   101,
		Name: "John",
	})

	repository.Add(Employee{
		ID:   102,
		Name: "Alice",
	})

	employee := repository.Get(101)

	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Employee Name:", employee.Name)
}