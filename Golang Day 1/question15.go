package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	ID   int
	Name string
	Role string
}

var employees []Employee
var currentID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Employee Management ---")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addEmployee(reader)
		case 2:
			searchEmployee()
		case 3:
			displayEmployees()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println("Exiting application...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addEmployee(reader *bufio.Reader) {
	fmt.Print("Enter employee name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter employee role: ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	emp := Employee{
		ID:   currentID,
		Name: name,
		Role: role,
	}
	employees = append(employees, emp)
	currentID++
	
	fmt.Println("Employee added successfully!")
}

func searchEmployee() {
	var id int
	fmt.Print("Enter employee ID to search: ")
	fmt.Scanln(&id)

	for _, emp := range employees {
		if emp.ID == id {
			fmt.Printf("Found -> ID: %d | Name: %s | Role: %s\n", emp.ID, emp.Name, emp.Role)
			return
		}
	}
	
	fmt.Println("Employee not found.")
}

func displayEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees to display.")
		return
	}

	fmt.Println("\nID\tName\t\tRole")
	fmt.Println("-----------------------------------")
	for _, emp := range employees {
		fmt.Printf("%d\t%s\t\t%s\n", emp.ID, emp.Name, emp.Role)
	}
}

func deleteEmployee() {
	var id int
	fmt.Print("Enter employee ID to delete: ")
	fmt.Scanln(&id)

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully!")
			return
		}
	}
	
	fmt.Println("Employee not found.")
}