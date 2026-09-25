package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func main() {

	fmt.Println("\n***************************************************************************")
	fmt.Println("10. Implement employee CRUD using a slice.")
	fmt.Println("*****************************************************************************")

	employees := []Employee{}

	var count int
	fmt.Print("How many employees do you want to add? ")
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		var id int
		var name string

		fmt.Printf("\nEnter details for Employee %d:\n", i+1)
		fmt.Print("Enter ID: ")
		fmt.Scanln(&id)

		fmt.Print("Enter Name: ")
		fmt.Scanln(&name)

		employees = append(employees, Employee{ID: id, Name: name})
	}

	fmt.Println("\n--- Current Employees ---")
	for i := 0; i < len(employees); i++ {
		fmt.Println("ID:", employees[i].ID, "| Name:", employees[i].Name)
	}

	if len(employees) > 0 {
		var updateID int
		var newName string

		fmt.Println("\n--- Update Employee ---")
		fmt.Print("Enter Employee ID to update: ")
		fmt.Scanln(&updateID)

		for i := 0; i < len(employees); i++ {
			if employees[i].ID == updateID {
				fmt.Print("Enter New Name: ")
				fmt.Scanln(&newName)

				employees[i].Name = newName
				fmt.Println("Name updated successfully!")
				break
			}
		}
	}

	if len(employees) > 0 {
		var deleteID int

		fmt.Println("\n--- Delete Employee ---")
		fmt.Print("Enter Employee ID to delete: ")
		fmt.Scanln(&deleteID)

		for i := 0; i < len(employees); i++ {
			if employees[i].ID == deleteID {
				employees = append(employees[:i], employees[i+1:]...)
				fmt.Println("Employee deleted successfully!")
				break
			}
		}
	}

	fmt.Println("\n--- Final Employee Records ---")
	if len(employees) == 0 {
		fmt.Println("No employee records found.")
	}
	for i := 0; i < len(employees); i++ {
		fmt.Println("ID:", employees[i].ID, "| Name:", employees[i].Name)
	}
}