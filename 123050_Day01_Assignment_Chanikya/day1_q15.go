package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Role string
}

func main() {
	var employees []Employee
	nextID := 1

	for {
		fmt.Println("\n--- EMPLOYEE MENU ---")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display All Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			var name string
			var role string

			fmt.Print("Enter Name (no spaces): ")
			fmt.Scanln(&name)
			fmt.Print("Enter Role (no spaces): ")
			fmt.Scanln(&role)

			newEmp := Employee{
				ID:   nextID,
				Name: name,
				Role: role,
			}

			employees = append(employees, newEmp)
			fmt.Println("Success! Added with ID:", nextID)
			nextID++

		} else if choice == 2 {
			var searchName string
			fmt.Print("Enter name to search: ")
			fmt.Scanln(&searchName)

			found := false
			for i := 0; i < len(employees); i++ {
				if employees[i].Name == searchName {
					fmt.Println("Found -> ID:", employees[i].ID, "| Name:", employees[i].Name, "| Role:", employees[i].Role)
					found = true
				}
			}
			if !found {
				fmt.Println("No employee found with that name.")
			}

		} else if choice == 3 {
			if len(employees) == 0 {
				fmt.Println("The list is empty.")
			} else {
				fmt.Println("\n--- ALL EMPLOYEES ---")
				for i := 0; i < len(employees); i++ {
					fmt.Println("ID:", employees[i].ID, "| Name:", employees[i].Name, "| Role:", employees[i].Role)
				}
			}

		} else if choice == 4 {
			var deleteID int
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scanln(&deleteID)

			var updatedList []Employee
			found := false

			for i := 0; i < len(employees); i++ {
				if employees[i].ID == deleteID {
					found = true
				} else {
					updatedList = append(updatedList, employees[i])
				}
			}

			employees = updatedList

			if found {
				fmt.Println("Employee deleted successfully.")
			} else {
				fmt.Println("ID not found.")
			}

		} else if choice == 5 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
