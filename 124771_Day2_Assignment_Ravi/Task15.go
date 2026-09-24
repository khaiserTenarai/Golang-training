package main

import (
	"fmt"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var employees []Employee
var employeeMap = make(map[int]Employee)

func main() {

	for {
		fmt.Println("\n*** Employee Management System ***")
		fmt.Println("1. Add Employee")
		fmt.Println("2. List Employees")
		fmt.Println("3. Find Employee")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addEmployee()

		case 2:
			listEmployees()

		case 3:
			findEmployee()

		case 4:
			updateEmployee()

		case 5:
			deleteEmployee()

		case 6:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func addEmployee() {
	var employee Employee
	fmt.Print("Enter ID: ")
	fmt.Scan(&employee.ID)

	if _, exists := employeeMap[employee.ID]; exists {
		fmt.Println("Employee ID already exists!")
		return
	}

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)
	fmt.Print("Enter Department: ")
	fmt.Scan(&employee.Department)
	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	employees = append(employees, employee)

	employeeMap[employee.ID] = employee

	fmt.Println("Employee added successfully!")
}

func listEmployees() {

	if len(employees) == 0 {
		fmt.Println("No employees found!")
		return
	}

	fmt.Println("\n Employee List ")

	for _, employee := range employees {
		fmt.Printf(
			"ID: %d | Name: %s | Department: %s | Salary: %.2f\n",
			employee.ID,
			employee.Name,
			employee.Department,
			employee.Salary,
		)
	}
}

func findEmployee() {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	employee, exists := employeeMap[id]

	if !exists {
		fmt.Println("Employee not found!")
		return
	}

	fmt.Println("\nEmployee Details")
	fmt.Println("----------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Salary:", employee.Salary)
}

func updateEmployee() {
	var id int
	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	employee, exists := employeeMap[id]

	if !exists {
		fmt.Println("Employee not found!")
		return
	}

	fmt.Print("Enter new Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter new Department: ")
	fmt.Scan(&employee.Department)

	fmt.Print("Enter new Salary: ")
	fmt.Scan(&employee.Salary)

	employeeMap[id] = employee

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees[i] = employee
			break
		}
	}

	fmt.Println("Employee updated successfully!")
}

func deleteEmployee() {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	_, exists := employeeMap[id]

	if !exists {
		fmt.Println("Employee not found!")
		return
	}

	delete(employeeMap, id)
	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {

			employees = append(
				employees[:i],
				employees[i+1:]...,
			)

			break
		}
	}

	fmt.Println("Employee deleted successfully!")
}
