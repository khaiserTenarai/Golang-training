package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var employees []Employee

// Create
func createEmployee(employee Employee) {
	employees = append(employees, employee)
	fmt.Println("Employee added")
}

// Read
func readEmployees() {
	for _, employee := range employees {
		fmt.Println(employee)
	}
}

// Update
func updateEmployee(id int, salary float64) {
	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees[i].Salary = salary
			fmt.Println("Employee updated")
			return
		}
	}

	fmt.Println("Employee not found")
}

// Delete
func deleteEmployee(id int) {
	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted")
			return
		}
	}

	fmt.Println("Employee not found")
}

func main() {

	createEmployee(Employee{1, "Rahul", 30000})
	createEmployee(Employee{2, "Priya", 40000})

	fmt.Println("Employees:")
	readEmployees()

	updateEmployee(1, 35000)

	fmt.Println("After Update:")
	readEmployees()

	deleteEmployee(2)

	fmt.Println("After Delete:")
	readEmployees()
}