package main

import (
	"fmt"
	"os"
)

type Employee struct {
	Name       string
	Department string
	Role       string
}

func main() {
	employeeDatabase := map[string]Employee{
		"101": {Name: "Alice Smith", Department: "Engineering", Role: "Developer"},
		"102": {Name: "Bob Jones", Department: "Design", Role: "UI Designer"},
		"103": {Name: "Charlie Brown", Department: "Marketing", Role: "Manager"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide an employee ID.")
		fmt.Println("Usage: go run main.go <employee_id>")
		return
	}

	targetID := os.Args[1]

	empInfo, exists := employeeDatabase[targetID]

	if exists {
		fmt.Printf("ID:         %s\n", targetID)
		fmt.Printf("Name:       %s\n", empInfo.Name)
		fmt.Printf("Department: %s\n", empInfo.Department)
		fmt.Printf("Role:       %s\n", empInfo.Role)
	} else {
		fmt.Printf("\nError: Employee with ID '%s' not found.\n\n", targetID)
	}
}
