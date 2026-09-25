package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

var employees []Employee

func addEmployee() {
	var id int
	var name string

	fmt.Print("Enter ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	employees = append(employees, Employee{ID: id, Name: name})
	fmt.Println("Employee added!")
}

func searchEmployee() {
	var id int
	fmt.Print("Enter ID to search: ")
	fmt.Scanln(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			fmt.Println("Found:", employees[i].ID, employees[i].Name)
			return
		}
	}
	fmt.Println("Employee not found.")
}

func displayEmployees() {
	fmt.Println("--- All Employees ---")
	if len(employees) == 0 {
		fmt.Println("No records found.")
		return
	}
	for i := 0; i < len(employees); i++ {
		fmt.Println("ID:", employees[i].ID, "| Name:", employees[i].Name)
	}
}

func deleteEmployee() {
	var id int
	fmt.Print("Enter ID to delete: ")
	fmt.Scanln(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted!")
			return
		}
	}
	fmt.Println("Employee not found.")
}

func main() {
	for {
		fmt.Println("\n1. Add | 2. Search | 3. Display | 4. Delete | 5. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:	
			addEmployee()
		case 2:
			searchEmployee()
		case 3:
			displayEmployees()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}