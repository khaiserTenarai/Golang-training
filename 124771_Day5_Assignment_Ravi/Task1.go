package main

import "fmt"

type Employee struct {
	ID           int
	Name         string
	Email        string
	Age          int
	Department   string
	Position     string
	Salary       float64
	MobileNumber string
}

func main() {
	employee := Employee{
		ID:           101,
		Name:         "Ravi Ranjan",
		Email:        "raviranjan.com",
		Age:          23,
		Department:   "Engineering",
		Position:     "Software Developer",
		Salary:       75000,
		MobileNumber: "9876543210",
	}

	fmt.Println(employee)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Department:", employee.Department)
}
