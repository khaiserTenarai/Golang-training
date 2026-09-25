package view

import (
	"fmt"
	"task15/task15/controller"
)

func Start() {
	var choice int
	var continueChoice string

	for {
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			controller.AddEmployee()
		case 2:
			controller.SearchEmployee()
		case 3:
			controller.DisplayEmployees()
		case 4:
			controller.DeleteEmployee()
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("Do you want to continue? Yes/No: ")
		fmt.Scan(&continueChoice)
		if continueChoice != "Yes" {
			break
		}
	}
}
