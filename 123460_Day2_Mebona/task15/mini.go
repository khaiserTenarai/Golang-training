package main

import (
	"fmt"
)

type Employee struct {
	ID   int
	Name string
	Role string
}

var employeeDB = make(map[int]Employee)
var orderedIDs []int

func main() {
	for {
		fmt.Println("\n=== Employee Management CLI ===")
		fmt.Println("1. Add New Employee")
		fmt.Println("2. View All Employees")
		fmt.Println("3. Update Existing Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Select an option (1-5): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addEmployee()
		case 2:
			viewEmployees()
		case 3:
			updateEmployee()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println("Exiting system. Goodbye!")
			return
		default:
			fmt.Println("Invalid selection. Please enter a number between 1 and 5.")
		}
	}
}

func addEmployee() {
	var id int
	var name, role string

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	if _, exists := employeeDB[id]; exists {
		fmt.Printf("Error: Employee with ID %d already exists.\n", id)
		return
	}

	fmt.Print("Enter First Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Role (single word): ")
	fmt.Scan(&role)

	employeeDB[id] = Employee{ID: id, Name: name, Role: role}
	orderedIDs = append(orderedIDs, id)

	fmt.Println("Success: Employee added.")
}

func viewEmployees() {
	if len(orderedIDs) == 0 {
		fmt.Println("The database is currently empty.")
		return
	}

	fmt.Println("\n--- Employee Directory ---")
	for _, id := range orderedIDs {
		emp := employeeDB[id]
		fmt.Printf("ID: %d | Name: %s | Role: %s\n", emp.ID, emp.Name, emp.Role)
	}
	fmt.Println("--------------------------")
}

func updateEmployee() {
	var id int
	fmt.Print("Enter Employee ID to update: ")
	fmt.Scan(&id)

	emp, exists := employeeDB[id]
	if !exists {
		fmt.Printf("Error: Employee with ID %d not found.\n", id)
		return
	}

	var newName, newRole string
	fmt.Printf("Enter New Name (current: %s): ", emp.Name)
	fmt.Scan(&newName)
	fmt.Printf("Enter New Role (current: %s): ", emp.Role)
	fmt.Scan(&newRole)

	emp.Name = newName
	emp.Role = newRole
	employeeDB[id] = emp

	fmt.Println("Success: Employee updated.")
}

func deleteEmployee() {
	var id int
	fmt.Print("Enter Employee ID to delete: ")
	fmt.Scan(&id)

	if _, exists := employeeDB[id]; !exists {
		fmt.Printf("Error: Employee with ID %d not found.\n", id)
		return
	}

	delete(employeeDB, id)

	for i, sliceID := range orderedIDs {
		if sliceID == id {
			orderedIDs = append(orderedIDs[:i], orderedIDs[i+1:]...)
			break
		}
	}

	fmt.Println("Success: Employee deleted.")
}