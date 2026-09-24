package main

import (
	employee "employeeapp/emp"
	"fmt"
)

func main() {
	// Creating an Employee struct directly
	emp1 := employee.Employee{
		ID:   101,
		Name: "Alice Smith",
		Role: "Engineering Manager",
	}

	// Using constructor function
	emp2 := employee.New(102, "Bob Jones")

	fmt.Println("--- Employee Directory ---")
	emp1.DisplayInfo()
	emp2.DisplayInfo()
}
