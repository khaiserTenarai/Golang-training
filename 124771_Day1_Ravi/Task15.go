package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {
	var employees []Employee
	var choice int

	for {
		fmt.Println("*** Employee Management System ***")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var emp Employee
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&emp.ID)

			fmt.Print("Enter Employee Name: ")
			fmt.Scan(&emp.Name)

			fmt.Print("Enter Employee Salary: ")
			fmt.Scan(&emp.Salary)

			employees = append(employees, emp)

			fmt.Println("Employee added successfully!")

		case 2:
			var id int
			found := false

			fmt.Print("Enter Employee ID to search: ")
			fmt.Scan(&id)

			for _, emp := range employees {
				if emp.ID == id {
					fmt.Println("\nEmployee Found!")
					fmt.Println("ID:", emp.ID)
					fmt.Println("Name:", emp.Name)
					fmt.Println("Salary:", emp.Salary)
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found!")
			}

		case 3:
			if len(employees) == 0 {
				fmt.Println("No employees available.")
			} else {
				fmt.Println("\n*** Employee List ***")

				for _, emp := range employees {
					fmt.Println("ID:", emp.ID)
					fmt.Println("Name:", emp.Name)
					fmt.Println("Salary:", emp.Salary)

				}
			}

		case 4:
			var id int
			found := false

			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&id)

			for i, emp := range employees {
				if emp.ID == id {
					employees = append(employees[:i], employees[i+1:]...)
					fmt.Println("Employee deleted successfully!")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found!")
			}

		case 5:
			fmt.Println("Exiting Employee Management System...")
			return

		default:
			fmt.Println("Invalid choice! Please enter a number from 1 to 5.")
		}
	}
}
