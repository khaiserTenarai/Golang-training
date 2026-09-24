package main

import (
	"fmt"
)

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func main() {
	
	employeeDB := make(map[int]Employee)

	
	employeeDB[101] = Employee{ID: 101, Name: "Alice", Position: "Developer", Salary: 75000}
	employeeDB[102] = Employee{ID: 102, Name: "Bob", Position: "Manager", Salary: 85000}
	employeeDB[103] = Employee{ID: 103, Name: "Charlie", Position: "Designer", Salary: 65000}

	fmt.Println("Database populated with 3 employees.\n")

	
	searchID := 102
	fmt.Printf(" Searching for Employee ID: %d ---\n", searchID)
	
	
	emp, exists := employeeDB[searchID]
	
	if exists {
		fmt.Printf("Success! Found %s working as a %s.\n", emp.Name, emp.Position)
	} else {
		fmt.Println("Error: Employee not found.")
	}

	fmt.Println("\nSearching for Employee ID: 999 ---")
	
	missingEmp, exists := employeeDB[999]
	
	if exists {
		fmt.Printf("Success! Found %s.\n", missingEmp.Name)
	} else {
		fmt.Printf("Error: No employee found with ID 999.\n")
	}
}