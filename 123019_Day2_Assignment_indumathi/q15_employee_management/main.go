package main

import (
	"fmt"

	"day2_employee_management/employee"
)

/*
Question 15:
Day 2 Mini Project: Build an in-memory Employee Management CLI
using slices, maps, loops, conditions and functions.
*/

func main() {
	for {
		fmt.Println(">>Employee Management System<<")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			employee.AddEmployee()

		case 2:
			employee.SearchEmployee()

		case 3:
			employee.DisplayEmployees()

		case 4:
			employee.DeleteEmployee()

		case 5:
			fmt.Println("Exiting Employee Management System.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
