package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

var employees []Employee

func addEmployee() {
	var employee Employee

	fmt.Println("\nAdd Employee")

	fmt.Print("Enter employee ID: ")
	fmt.Scanln(&employee.ID)

	fmt.Print("Enter employee name: ")
	fmt.Scanln(&employee.Name)

	fmt.Print("Enter employee age: ")
	fmt.Scanln(&employee.Age)

	fmt.Print("Enter employee salary: ")
	fmt.Scanln(&employee.Salary)

	employees = append(employees, employee)

	fmt.Println("Employee added successfully.")
}

func searchEmployee() {
	var id int

	fmt.Println("\nSearch Employee")
	fmt.Print("Enter employee ID: ")
	fmt.Scanln(&id)

	for _, employee := range employees {
		if employee.ID == id {
			fmt.Println("Employee found.")
			fmt.Println("ID:", employee.ID)
			fmt.Println("Name:", employee.Name)
			fmt.Println("Age:", employee.Age)
			fmt.Println("Salary:", employee.Salary)
			return
		}
	}

	fmt.Println("Employee not found.")
}

func displayEmployees() {
	fmt.Println("\nEmployee List")

	if len(employees) == 0 {
		fmt.Println("No employees available.")
		return
	}

	for _, employee := range employees {
		fmt.Println("--------------------")
		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Age:", employee.Age)
		fmt.Println("Salary:", employee.Salary)
	}
}

func deleteEmployee() {
	var id int

	fmt.Println("\nDelete Employee")
	fmt.Print("Enter employee ID: ")
	fmt.Scanln(&id)

	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully.")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func main() {
	for {
		fmt.Println("===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
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
			fmt.Println("Program ended.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
