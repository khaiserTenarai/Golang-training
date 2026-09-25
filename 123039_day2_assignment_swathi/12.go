
package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func main() {

	employees := []Employee{}

	for {

		fmt.Println("\n===== EMPLOYEE MENU =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			var employee Employee

			fmt.Println("\n===== ADD EMPLOYEE =====")

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&employee.ID)

			fmt.Print("Enter Employee Name: ")
			fmt.Scan(&employee.Name)

			employees = append(employees, employee)

			fmt.Println("Employee added successfully.")

		case 2:

			var id int

			fmt.Println("\n===== SEARCH EMPLOYEE =====")

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for i := 0; i < len(employees); i++ {

				if employees[i].ID == id {

					fmt.Println("Employee ID:", employees[i].ID)
					fmt.Println("Employee Name:", employees[i].Name)

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 3:

			fmt.Println("\n===== EMPLOYEE LIST =====")

			if len(employees) == 0 {
				fmt.Println("No employees found.")
			} else {

				for i := 0; i < len(employees); i++ {

					fmt.Println("--------------------")
					fmt.Println("Employee ID:", employees[i].ID)
					fmt.Println("Employee Name:", employees[i].Name)
				}
			}

		case 4:

			var id int

			fmt.Println("\n===== DELETE EMPLOYEE =====")

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for i := 0; i < len(employees); i++ {

				if employees[i].ID == id {

					employees = append(employees[:i], employees[i+1:]...)

					fmt.Println("Employee deleted successfully.")

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 5:

			fmt.Println("Exiting program...")
			return

		default:

			fmt.Println("Invalid choice. Please enter 1 to 5.")
		}
	}
}

