package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Employee struct {
	ID       int
	Name     string
	Email    string
	Position string
}

var employees = []Employee{
	{
		ID:       1,
		Name:     "Ravi",
		Email:    "ravi@gmail.com",
		Position: "Developer",
	},
	{
		ID:       2,
		Name:     "Amit",
		Email:    "amit@gmail.com",
		Position: "Tester",
	},
}

var nextID = 3

func CreateEmployee(name, email, position string) {
	employee := Employee{
		ID:       nextID,
		Name:     name,
		Email:    email,
		Position: position,
	}

	employees = append(employees, employee)
	nextID++

	fmt.Println("Employee created successfully!")
}

func GetEmployee() {
	fmt.Println("\n===== Employees =====")

	for _, employee := range employees {
		fmt.Println("----------------------")
		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Email:", employee.Email)
		fmt.Println("Position:", employee.Position)
	}
}

func GetEmployee(id int) {
	for _, employee := range employees {
		if employee.ID == id {
			fmt.Println("\nID:", employee.ID)
			fmt.Println("Name:", employee.Name)
			fmt.Println("Email:", employee.Email)
			fmt.Println("Position:", employee.Position)
			return
		}
	}

	fmt.Println("Employee not found.")
}

func UpdateEmployee(id int, name, email, position string) {
	for i := range employees {
		if employees[i].ID == id {
			employees[i].Name = name
			employees[i].Email = email
			employees[i].Position = position

			fmt.Println("Employee updated successfully!")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func DeleteEmployee(id int) {
	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully!")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func GetEmployeeID(value string) (int, bool) {
	value = strings.TrimSpace(value)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, false
	}

	return id, true
}

func main() {

	for {
		fmt.Println("\n======================")
		fmt.Println("    EMPLOYEE CRUD")
		fmt.Println("======================")
		fmt.Println("1. Create Employee")
		fmt.Println("2. Get All Employees")
		fmt.Println("3. Get Employee By ID")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")
		fmt.Println("======================")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var name, email, position string

			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			fmt.Print("Enter email: ")
			fmt.Scan(&email)

			fmt.Print("Enter position: ")
			fmt.Scan(&position)

			CreateEmployee(name, email, position)

		case 2:

			GetEmployee(id)

		case 3:
			var value string

			fmt.Print("Enter employee ID: ")
			fmt.Scan(&value)

			id, valid := GetEmployeeID(value)

			if !valid {
				fmt.Println("Invalid ID.")
				continue
			}

			GetEmployee(id)

		case 4:
			var value string

			fmt.Print("Enter employee ID: ")
			fmt.Scan(&value)

			id, valid := GetEmployeeID(value)

			if !valid {
				fmt.Println("Invalid ID.")
				continue
			}

			var name, email, position string

			fmt.Print("Enter new name: ")
			fmt.Scan(&name)

			fmt.Print("Enter new email: ")
			fmt.Scan(&email)

			fmt.Print("Enter new position: ")
			fmt.Scan(&position)

			UpdateEmployee(id, name, email, position)

		case 5:
			var value string

			fmt.Print("Enter employee ID: ")
			fmt.Scan(&value)

			id, valid := GetEmployeeID(value)

			if !valid {
				fmt.Println("Invalid ID.")
				continue
			}

			DeleteEmployee(id)

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

