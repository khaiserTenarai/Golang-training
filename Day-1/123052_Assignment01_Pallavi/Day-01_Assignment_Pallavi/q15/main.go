package main

import "fmt"

// Employee represents an employee
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	// Slice to store employees
	var employees []Employee

	var choice int

	for {

		fmt.Println("\n******** Employee Management ********")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		// Add Employee
		case 1:

			var id int
			var name string
			var salary float64

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter Employee Name: ")
			fmt.Scan(&name)

			fmt.Print("Enter Employee Salary: ")
			fmt.Scan(&salary)

			employee := Employee{
				ID:     id,
				Name:   name,
				Salary: salary,
			}

			employees = append(employees, employee)

			fmt.Println("Employee added successfully!")

		// Search Employee
		case 2:

			var id int
			found := false

			fmt.Print("Enter Employee ID to search: ")
			fmt.Scan(&id)

			for _, employee := range employees {

				if employee.ID == id {

					fmt.Println("\nEmployee Found")
					fmt.Println("ID:", employee.ID)
					fmt.Println("Name:", employee.Name)
					fmt.Println("Salary:", employee.Salary)

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found")
			}

		// Display Employees
		case 3:

			if len(employees) == 0 {

				fmt.Println("No employees available")

			} else {

				fmt.Println("\n******** Employee List ********")

				for _, employee := range employees {

					fmt.Println("ID:", employee.ID)
					fmt.Println("Name:", employee.Name)
					fmt.Println("Salary:", employee.Salary)
					fmt.Println("----------------------------")
				}
			}

		// Delete Employee
		case 4:

			var id int
			found := false

			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&id)

			for i, employee := range employees {

				if employee.ID == id {

					// Remove employee from slice
					employees = append(employees[:i], employees[i+1:]...)

					fmt.Println("Employee deleted successfully!")

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found")
			}

		// Exit
		case 5:

			fmt.Println("Thank you!")
			return

		default:

			fmt.Println("Invalid choice")
		}
	}
}