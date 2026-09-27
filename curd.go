package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	employees := []Employee{}

	for {
		fmt.Println("\n===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			// Add Employee
			var emp Employee

			fmt.Print("Enter ID: ")
			fmt.Scan(&emp.ID)

			fmt.Print("Enter Name: ")
			fmt.Scan(&emp.Name)

			fmt.Print("Enter Salary: ")
			fmt.Scan(&emp.Salary)

			employees = append(employees, emp)

			fmt.Println("Employee added successfully!")

		case 2:
			// Search Employee
			var id int

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for _, emp := range employees {
				if emp.ID == id {
					fmt.Println("Employee Found!")
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
			// Update Employee
			var id int

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for i := range employees {
				if employees[i].ID == id {

					fmt.Print("Enter New Name: ")
					fmt.Scan(&employees[i].Name)

					fmt.Print("Enter New Salary: ")
					fmt.Scan(&employees[i].Salary)

					fmt.Println("Employee updated successfully!")

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found!")
			}

		case 4:
			// Delete Employee
			var id int

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for i, emp := range employees {
				if emp.ID == id {

					employees = append(
						employees[:i],
						employees[i+1:]...,
					)

					fmt.Println("Employee deleted successfully!")

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found!")
			}

		case 5:
			// Exit
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
