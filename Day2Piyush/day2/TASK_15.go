package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var employees []Employee

func addEmployee() {
	var id int
	var name string
	var salary float64

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Employee Salary: ")
	fmt.Scan(&salary)

	employee := Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}

	employees = append(employees, employee)

	fmt.Println("Employee added successfully!")
}

func viewEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println("\n******** Employee Lit********")

	for _, employee := range employees {
		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Salary:", employee.Salary)
		fmt.Println("--------------------")
	}
}

func searchEmployee() {
	var id int

	fmt.Print("Enter Employee ID to search: ")
	fmt.Scan(&id)

	for _, employee := range employees {
		if employee.ID == id {
			fmt.Println("\nEmployee Found!")
			fmt.Println("ID:", employee.ID)
			fmt.Println("Name:", employee.Name)
			fmt.Println("Salary:", employee.Salary)
			return
		}
	}

	fmt.Println("Employee not found.")
}

func updateEmployee() {
	var id int

	fmt.Print("Enter Employee ID to update: ")
	fmt.Scan(&id)

	for i := range employees {
		if employees[i].ID == id {

			fmt.Print("Enter new name: ")
			fmt.Scan(&employees[i].Name)

			fmt.Print("Enter new salary: ")
			fmt.Scan(&employees[i].Salary)

			fmt.Println("Employee updated successfully!")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func deleteEmployee() {
	var id int

	fmt.Print("Enter Employee ID to delete: ")
	fmt.Scan(&id)

	for i := range employees {
		if employees[i].ID == id {

			employees = append(employees[:i], employees[i+1:]...)

			fmt.Println("Employee deleted successfully!")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func main() {

	for {

		fmt.Println("\n===== Employee Management System =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Search Employee")
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
			viewEmployees()

		case 3:
			searchEmployee()

		case 4:
			updateEmployee()

		case 5:
			deleteEmployee()

		case 6:
			fmt.Println("Program exited.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}