package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID     int
	Name   string
	Role   string
	Salary int
}

func main() {
	// Employee data stored in memory
	employees := []Employee{
		{ID: 1, Name: "Rahul", Role: "Developer", Salary: 50000},
		{ID: 2, Name: "Priya", Role: "Designer", Salary: 45000},
		{ID: 3, Name: "Amit", Role: "Manager", Salary: 70000},
	}

	// Check if employee ID was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <employee-id>")
		return
	}

	// Convert command-line argument to integer
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid employee ID")
		return
	}

	// Search employee
	for _, employee := range employees {
		if employee.ID == id {
			fmt.Println("Employee found:")
			fmt.Println("ID:", employee.ID)
			fmt.Println("Name:", employee.Name)
			fmt.Println("Role:", employee.Role)
			fmt.Println("Salary:", employee.Salary)
			return
		}
	}

	fmt.Println("Employee not found")
}

/*
Execution & Output :
---------------------------

PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy> go run .\11-CLI-employee-search.go 1
Employee found:
ID: 1
Name: Rahul
Role: Developer
Salary: 50000

PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy> go run .\11-CLI-employee-search.go 2
Employee found:
ID: 2
Name: Priya
Role: Designer
Salary: 45000

PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy> go run .\11-CLI-employee-search.go 34
Employee not found
*/
