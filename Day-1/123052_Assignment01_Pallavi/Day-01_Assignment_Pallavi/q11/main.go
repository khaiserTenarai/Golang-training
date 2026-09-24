package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	employees := map[int]Employee{
		101: {ID: 101, Name: "Pallavi", Salary: 50000},
		102: {ID: 102, Name: "Rahul", Salary: 45000},
		103: {ID: 103, Name: "Anita", Salary: 55000},
	}

	// Check whether employee ID was provided
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
	employee, found := employees[id]

	if found {
		fmt.Println("Employee ID:", employee.ID)
		fmt.Println("Employee Name:", employee.Name)
		fmt.Println("Employee Salary:", employee.Salary)
	} else {
		fmt.Println("Employee not found")
	}
}