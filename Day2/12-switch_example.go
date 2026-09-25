package main

import "fmt"

func main() {

	var choice int

	fmt.Println("===== Employee Menu =====")
	fmt.Println("1. Add Employee")
	fmt.Println("2. View Employee")
	fmt.Println("3. Update Employee")
	fmt.Println("4. Delete Employee")
	fmt.Println("5. Exit")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	// Here In Go Language , we no need to use break, it will automatically stop execution after particular case
	switch choice {

	case 1:
		fmt.Println("Add Employee selected")

	case 2:
		fmt.Println("View Employee selected")

	case 3:
		fmt.Println("Update Employee selected")

	case 4:
		fmt.Println("Delete Employee selected")

	case 5:
		fmt.Println("Exit selected")

	default:
		fmt.Println("Invalid choice")
	}
}
