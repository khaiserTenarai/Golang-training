package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var employees = []Employee{}

func createEmployee() {
	var id int
	var name string
	var salary float64

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&salary)

	employee := Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}

	employees = append(employees, employee)

	fmt.Println("Employee created successfully")
}

func viewEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}

	for _, employee := range employees {
		fmt.Println("----------------------")
		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Salary:", employee.Salary)
	}
}

func findEmployee() {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	for _, employee := range employees {
		if employee.ID == id {
			fmt.Println("Employee Found")
			fmt.Println("Name:", employee.Name)
			fmt.Println("Salary:", employee.Salary)
			return
		}
	}

	fmt.Println("Employee not found")
}

func deleteEmployee() {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)

			fmt.Println("Employee deleted")
			return
		}
	}

	fmt.Println("Employee not found")
}

func main() {

	for {
		fmt.Println()
		fmt.Println("******** Employee Management ********")
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