package main

import (
	"fmt"
	"os"
)

type Employee struct {
	ID   string
	Name string
	Dept string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an employee ID.")
		return
	}

	searchID := os.Args[1]

	employees := []Employee{
		{"101", "Sachin", "Engineering"},
		{"102", "Nakul", "HR"},
		{"103", "Bhavana", "Sales"},
	}

	for _, emp := range employees {
		if emp.ID == searchID {
			fmt.Printf("ID   : %s\nName : %s\nDept : %s\n", emp.ID, emp.Name, emp.Dept)
			return
		}
	}

	fmt.Println("No employee found with ID:", searchID)
}
