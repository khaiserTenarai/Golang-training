package main

import "fmt"

// Employee represents an employee in our application.
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {

	// Create an Employee object.
	employee := Employee{
		ID:         101,
		Name:       "Alice",
		Department: "Engineering",
		Salary:     75000,
	}

	// Display employee information.
	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Salary:", employee.Salary)
}

// gofmt → formats Go source code.

// -w → writes the formatted code back to the file.

// gofmt -w main.go
