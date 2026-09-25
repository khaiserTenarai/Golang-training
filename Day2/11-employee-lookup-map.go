package main

import "fmt"

// Employee structure
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	// Create a map where ID is the key
	employees := make(map[int]Employee)

	var choice int

	// Keep showing menu until user chooses Exit
	for {
		fmt.Println("\n--- Employee Lookup ---")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		// Add employee to map
		case 1:
			var id int
			var name string
			var salary float64

			fmt.Print("Enter ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter Name: ")
			fmt.Scan(&name)

			fmt.Print("Enter Salary: ")
			fmt.Scan(&salary)

			// Store employee using ID as key
			employees[id] = Employee{id, name, salary}

			fmt.Println("Employee added successfully.")

		// Search employee using ID
		case 2:
			var id int

			fmt.Print("Enter Employee ID to search: ")
			fmt.Scan(&id)

			// Get employee from map
			employee, found := employees[id]

			if found {
				fmt.Println("\nEmployee Found")
				fmt.Println("ID:", employee.ID)
				fmt.Println("Name:", employee.Name)
				fmt.Println("Salary:", employee.Salary)
			} else {
				fmt.Println("Employee not found.")
			}

		// Display all employees
		case 3:
			fmt.Println("\n--- Employee List ---")

			// Loop through the map
			for _, employee := range employees {
				fmt.Println("ID:", employee.ID)
				fmt.Println("Name:", employee.Name)
				fmt.Println("Salary:", employee.Salary)
				fmt.Println()
			}

		// Exit the program
		case 4:
			fmt.Println("Program ended.")
			return

		// Handle invalid choice
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
