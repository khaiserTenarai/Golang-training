package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary float64
}

func createEmployee(employees []Employee, employee Employee) []Employee {
	employees = append(employees, employee)
	return employees
}

func readEmployees(employees []Employee) {
	for _, employee := range employees {
		fmt.Println(employee)
	}
}

func updateEmployee(employees []Employee, id int, salary float64) {
	for i := range employees {
		if employees[i].id == id {
			employees[i].salary = salary
		}
	}
}

func deleteEmployee(employees []Employee, id int) []Employee {
	for i := range employees {
		if employees[i].id == id {
			employees = append(employees[:i], employees[i+1:]...)
		}
	}
	return employees
}

func main() {
	employees := []Employee{}

	employees = createEmployee(employees, Employee{101, "Indu", 50000})
	employees = createEmployee(employees, Employee{102, "Arun", 60000})

	fmt.Println("Employees:")
	readEmployees(employees)

	updateEmployee(employees, 101, 55000)

	fmt.Println("After update:")
	readEmployees(employees)

	employees = deleteEmployee(employees, 102)

	fmt.Println("After delete:")
	readEmployees(employees)
}