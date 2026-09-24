package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
	Email  string
}

func main() {
	fmt.Println("Employee Management Application")
}

func createEmployee() Employee {
	return Employee{
		ID:     1,
		Name:   "John",
		Salary: 50000,
		Email:  "john@example.com",
	}
}

func findEmployee(employees map[int]Employee, id int) {
	employee, exists := employees[id]

	if exists {
		fmt.Println("Employee:", employee.Name)
	} else {
		fmt.Println("Employee not found")
	}
}
func menu() {
	fmt.Println("1. Create Employee")
	fmt.Println("2. Find Employee")
	fmt.Println("3. Exit")
}