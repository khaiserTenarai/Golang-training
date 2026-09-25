package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var employees []Employee
var employeeMap = make(map[int]string)

func createEmployee() {
	var id int
	var name string
	var salary float64

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	existingName, exists := employeeMap[id]
	if exists {
		fmt.Println("Error: ID already exists for employee:", existingName)
		return
	}

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&salary)

	newEmp := Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}

	employees = append(employees, newEmp)

	employeeMap[id] = name

	fmt.Println("Employee created successfully!")
}

func viewEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for i := 0; i < len(employees); i++ {
		fmt.Println("----------------------")
		fmt.Println("ID    :", employees[i].ID)
		fmt.Println("Name  :", employees[i].Name)
		fmt.Println("Salary:", employees[i].Salary)
	}
}

func findEmployee() {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	name, exists := employeeMap[id]
	if !exists {
		fmt.Println("Employee not found.")
		return
	}

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			fmt.Println("--- Employee Found ---")
			fmt.Println("ID    :", employees[i].ID)
			fmt.Println("Name  :", name)
			fmt.Println("Salary:", employees[i].Salary)
			return
		}
	}
}

func deleteEmployee() {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	_, exists := employeeMap[id]
	if !exists {
		fmt.Println("Employee not found.")
		return
	}

	delete(employeeMap, id)

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully!")
			return
		}
	}
}

func main() {
	for {
		fmt.Println("\n")
		fmt.Println("1. Create Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Find Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createEmployee()
		case 2:
			viewEmployees()
		case 3:
			findEmployee()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}