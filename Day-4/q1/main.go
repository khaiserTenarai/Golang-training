package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Create
func createEmployee(id int, name string, salary float64) Employee {
	return Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}
}

// Read
func readEmployee(employee Employee) {
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Salary:", employee.Salary)
}

// Update
func updateEmployee(employee *Employee, name string, salary float64) {
	employee.Name = name
	employee.Salary = salary
}

// Delete
func deleteEmployee(employee *Employee) {
	employee.ID = 0
	employee.Name = ""
	employee.Salary = 0
}

func main() {

	employee := createEmployee(101, "Pallavi", 50000)

	fmt.Println("Employee Created:")
	readEmployee(employee)

	updateEmployee(&employee, "Pallavi Updated", 55000)

	fmt.Println("\nAfter Update:")
	readEmployee(employee)

	deleteEmployee(&employee)

	fmt.Println("\nAfter Delete:")
	readEmployee(employee)
}