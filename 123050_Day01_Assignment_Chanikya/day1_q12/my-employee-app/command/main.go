// command/main.go
package main

import (
	"my-employee-app/internal/employee"
)

func main() {
	// Create a sample employee instance
	emp1 := employee.Employee{
		ID:       101,
		Name:     "Alice Smith",
		Position: "Software Engineer",
		Salary:   85000.00,
	}

	// Call the internal package function to display info
	employee.Display(emp1)
}
