package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID   int
	Name string
	Role string
}

var employees []Employee
var nextID = 1

func createEmployee(name string, role string) {
	emp := Employee{
		ID:   nextID,
		Name: name,
		Role: role,
	}
	employees = append(employees, emp)
	nextID++
}

func readAllEmployees() []Employee {
	return employees
}

func readEmployeeByID(id int) (Employee, error) {
	for _, emp := range employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

func updateEmployee(id int, newName string, newRole string) error {
	for i, emp := range employees {
		if emp.ID == id {
			employees[i].Name = newName
			employees[i].Role = newRole
			return nil
		}
	}
	return errors.New("employee not found")
}

func deleteEmployee(id int) error {
	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return errors.New("employee not found")
}

func main() {
	createEmployee("Lakshmi", "Software Engineer")
	createEmployee("Ayush", "Product Manager")

	fmt.Println(readAllEmployees())

	emp, err := readEmployeeByID(1)
	if err == nil {
		fmt.Println(emp)
	}

	updateEmployee(2, "Riya", "Director")

	fmt.Println(readAllEmployees())

	deleteEmployee(1)

	fmt.Println(readAllEmployees())
}