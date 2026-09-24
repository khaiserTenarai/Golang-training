package main

import "fmt"

// Employee structure to hold record data
type Employee struct {
	ID         string
	Name       string
	Department string
}

// In-memory map to store employees: key is ID, value is Employee
var employees = map[string]Employee{
	"101": {ID: "101", Name: "Alice", Department: "IT"},
	"102": {ID: "102", Name: "Bob", Department: "HR"},
}

// Add a new employee
func AddEmployee(id, name, dept string) {
	if _, exists := employees[id]; exists {
		fmt.Println("\nError: Employee ID already exists!")
		return
	}

	employees[id] = Employee{
		ID:         id,
		Name:       name,
		Department: dept,
	}
	fmt.Println("\nEmployee added successfully!")
}

// Search for an employee by ID
func SearchEmployee(id string) {
	emp, exists := employees[id]
	if !exists {
		fmt.Println("\nError: Employee not found!")
		return
	}

	fmt.Println("\n--- Employee Found ---")
	fmt.Printf("ID:         %s\n", emp.ID)
	fmt.Printf("Name:       %s\n", emp.Name)
	fmt.Printf("Department: %s\n", emp.Department)
}

// Display all stored employees
func DisplayAllEmployees() {
	if len(employees) == 0 {
		fmt.Println("\nNo employees found.")
		return
	}

	fmt.Println("\n--- All Employees ---")
	for _, emp := range employees {
		fmt.Printf("ID: %s | Name: %s | Dept: %s\n", emp.ID, emp.Name, emp.Department)
	}
}

// Delete an employee by ID
func DeleteEmployee(id string) {
	if _, exists := employees[id]; !exists {
		fmt.Println("\nError: Employee not found!")
		return
	}

	delete(employees, id)
	fmt.Println("\nEmployee deleted successfully!")
}
