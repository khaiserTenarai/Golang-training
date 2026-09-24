package main

import "fmt"

func main() {
	for {
		ShowMenu()
		choice := ReadInput("Enter choice (1-5): ")

		switch choice {
		case "1":
			id := ReadInput("Enter Employee ID: ")
			name := ReadInput("Enter Name: ")
			dept := ReadInput("Enter Department: ")
			AddEmployee(id, name, dept)

		case "2":
			id := ReadInput("Enter Employee ID to search: ")
			SearchEmployee(id)

		case "3":
			DisplayAllEmployees()

		case "4":
			id := ReadInput("Enter Employee ID to delete: ")
			DeleteEmployee(id)

		case "5":
			fmt.Println("\nGoodbye!")
			return

		default:
			fmt.Println("\nInvalid choice! Please select 1 to 5.")
		}
	}
}
