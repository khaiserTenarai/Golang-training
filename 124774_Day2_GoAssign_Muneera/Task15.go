package main

import (
	"fmt"
)

type Employee struct {
	id     int
	name   string
	salary int
}

var employees []Employee

func addEmployee() {
	var id int
	var name string
	var salary int

	fmt.Println("Enter id:")
	fmt.Scan(&id)

	fmt.Println("Enter name:")
	fmt.Scan(&name)

	fmt.Println("Enter salary:")
	fmt.Scan(&salary)

	employees = append(employees, Employee{id, name, salary})

	fmt.Println("Employee added successfully!")
}

func showEmployee() {
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for i := 0; i < len(employees); i++ {
		fmt.Println(
			"ID:", employees[i].id,
			"Name:", employees[i].name,
			"Salary:", employees[i].salary,
		)
	}
}

func searchEmployee() {
	var id int
	fmt.Println("Enter employee id:")
	fmt.Scan(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].id == id {
			fmt.Println("Name:", employees[i].name)
			fmt.Println("Salary:", employees[i].salary)
			return
		}
	}
	fmt.Println("Employee not found")
}

func deleteEmployee() {
	var id int
	fmt.Println("Enter employee id:")
	fmt.Scan(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].id == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted")
			return
		}
	}
	fmt.Println("Employee not found")
}

func main() {
	for {
		fmt.Println("\n------Employee Management--------")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Show Employee")
		fmt.Println("3. Search Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Println("Enter Choice:")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if choice == 1 {
			addEmployee()
		} else if choice == 2 {
			showEmployee()
		} else if choice == 3 {
			searchEmployee()
		} else if choice == 4 {
			deleteEmployee()
		} else if choice == 5 {
			fmt.Println("Thank You")
			break
		} else {
			fmt.Println("Invalid Choice")
		}
	}
}
