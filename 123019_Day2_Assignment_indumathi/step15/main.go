package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var employees []Employee
var employeeMap = make(map[int]Employee)

func addEmployee() {
	var emp Employee

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&emp.ID)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&emp.Name)

	fmt.Print("Enter Department: ")
	fmt.Scan(&emp.Department)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&emp.Salary)

	employees = append(employees, emp)
	employeeMap[emp.ID] = emp

	fmt.Println("Employee added successfully.")
}

func viewEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, emp := range employees {
		fmt.Println("ID:", emp.ID)
		fmt.Println("Name:", emp.Name)
		fmt.Println("Department:", emp.Department)
		fmt.Println("Salary:", emp.Salary)
		fmt.Println()
	}
}

func searchEmployee() {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	emp, exists := employeeMap[id]

	if exists {
		fmt.Println("ID:", emp.ID)
		fmt.Println("Name:", emp.Name)
		fmt.Println("Department:", emp.Department)
		fmt.Println("Salary:", emp.Salary)
	} else {
		fmt.Println("Employee not found.")
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

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}

	fmt.Println("Employee deleted successfully.")
}

func main() {
	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Search Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addEmployee()

		case 2:
			viewEmployees()

		case 3:
			searchEmployee()

		case 4:
			deleteEmployee()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}