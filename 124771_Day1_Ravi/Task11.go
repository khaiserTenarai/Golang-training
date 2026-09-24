package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	searchID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Invalid employee ID. Please enter a valid number.")
		return
	}

	// In-memory employee database
	employees := map[int]Employee{
		101: {ID: 101, Name: "Ravi", Department: "Engineering", Salary: 85000.00},
		102: {ID: 102, Name: "Amit", Department: "Marketing", Salary: 62000.00},
		103: {ID: 103, Name: "John", Department: "Sales", Salary: 70000.00},
	}

	emp, found := employees[searchID]
	if !found {
		fmt.Printf("Employee with ID %d not found.\n", searchID)
		return
	}

	fmt.Println("--- Employee Found ---")
	fmt.Printf("ID: %d\n", emp.ID)
	fmt.Printf("Name: %s\n", emp.Name)
	fmt.Printf("Department: %s\n", emp.Department)
	fmt.Printf("Salary: $%.2f\n", emp.Salary)
}
