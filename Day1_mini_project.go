package main

import "fmt"

var employees []string

func addEmployee() {
	var name string

	fmt.Print("Enter employee name: ")
	fmt.Scan(&name)

	employees = append(employees, name)

	fmt.Println("Employee added successfully!")
}

func searchEmployee() {
	var name string

	fmt.Print("Enter employee name to search: ")
	fmt.Scan(&name)

	for _, employee := range employees {
		if employee == name {
			fmt.Println("Employee found!")
			return
		}
	}

	fmt.Println("Employee not found!")
}

func displayEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found!")
		return
	}

	fmt.Println("\nEmployee List:")

	for i, employee := range employees {
		fmt.Println(i+1, employee)
	}
}

func deleteEmployee() {
	var name string

	fmt.Print("Enter employee name to delete: ")
	fmt.Scan(&name)

	for i, employee := range employees {
		if employee == name {

			employees = append(employees[:i], employees[i+1:]...)

			fmt.Println("Employee deleted successfully!")
			return
		}
	}

	fmt.Println("Employee not found!")
}

func main() {

	for {
		fmt.Println("\n===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

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
