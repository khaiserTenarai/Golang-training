
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID    int
	Name  string
	Email string
	Role  string
}

func main() {

	employees := []Employee{
		{ID: 101, Name: "Swathi", Email: "swathi@example.com", Role: "Developer"},
		{ID: 102, Name: "Rahul", Email: "rahul@example.com", Role: "Tester"},
		{ID: 103, Name: "Priya", Email: "priya@example.com", Role: "Manager"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide employee ID")
		fmt.Println("Example: go run . 101")
		return
	}

	employeeID, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Invalid employee ID")
		return
	}

	for i := 0; i < len(employees); i++ {

		if employees[i].ID == employeeID {
			fmt.Println("===== EMPLOYEE DETAILS =====")
			fmt.Println("Employee ID:", employees[i].ID)
			fmt.Println("Name:", employees[i].Name)
			fmt.Println("Email:", employees[i].Email)
			fmt.Println("Role:", employees[i].Role)
			return
		}
	}

	fmt.Println("Employee not found")
}


