package main

import "fmt"

// Employee structure
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	// Create an empty slice of employees
	employees := []Employee{}

	var choice int

	// Keep showing the menu until user chooses Exit
	for {
		fmt.Println("\n--- Employee CRUD ---")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		// Add a new employee
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

			// Add employee to the slice
			employees = append(employees, Employee{id, name, salary})

			fmt.Println("Employee added successfully.")

		// Display all employees
		case 2:
			fmt.Println("\n--- Employee List ---")

			if len(employees) == 0 {
				fmt.Println("No employees found.")
			} else {

				// Loop through the slice
				for _, emp := range employees {
					fmt.Println("ID:", emp.ID)
					fmt.Println("Name:", emp.Name)
					fmt.Println("Salary:", emp.Salary)
					fmt.Println()
				}
			}

		// Update an employee
		case 3:
			var id int

			fmt.Print("Enter Employee ID to update: ")
			fmt.Scan(&id)

			found := false

			// Search employee using ID
			for i := range employees {
				if employees[i].ID == id {

					fmt.Print("Enter new name: ")
					fmt.Scan(&employees[i].Name)

					fmt.Print("Enter new salary: ")
					fmt.Scan(&employees[i].Salary)

					fmt.Println("Employee updated successfully.")

					found = true
					break
				}
			}

			// If ID was not found
			if !found {
				fmt.Println("Employee not found.")
			}

		// Delete an employee
		case 4:
			var id int

			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&id)

			found := false

			// Search employee using ID
			for i := range employees {
				if employees[i].ID == id {

					// Remove employee from slice
					employees = append(employees[:i], employees[i+1:]...)

					fmt.Println("Employee deleted successfully.")

					found = true
					break
				}
			}

			// If ID was not found
			if !found {
				fmt.Println("Employee not found.")
			}

		// Exit the program
		case 5:
			fmt.Println("Program ended.")
			return

		// Handle invalid menu choice
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
