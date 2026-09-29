package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

func main() {

	employees := []Employee{
		{1, "Pooja", 23, 50000},
		{2, "Rahul", 25, 60000},
		{3, "Amit", 27, 70000},
	}

	if len(os.Args) < 2 {
		fmt.Println("Please enter employee ID")
		fmt.Println("Example: go run main.go 2")
		return
	}

	// Get employee ID from command line
	id, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Invalid employee ID")
		return
	}

	// Search employee
	for i, employee := range employees {

		if employee.ID == id {

			fmt.Println("\nEmployee Found! at the index", i)
			fmt.Println("ID:", employee.ID)
			fmt.Println("Name:", employee.Name)
			fmt.Println("Age:", employee.Age)
			fmt.Println("Salary:", employee.Salary)

			return
		}
	}

	fmt.Println("Employee not found")
}
