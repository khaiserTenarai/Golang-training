package main

import "fmt"

func main() {
	for {
		var choice int

		fmt.Println("\nEmployee Menu")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Add Employee selected")

		case 2:
			fmt.Println("View Employees selected")

		case 3:
			fmt.Println("Update Employee selected")

		case 4:
			fmt.Println("Delete Employee selected")

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}