// 1. Create functions for employee CRUD.

package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

var employees []Employee

func addEmployee(id int, name string) {
	employees = append(employees, Employee{ID: id, Name: name})
}

func getEmployee(id int) (Employee, bool) {
	for _, emp := range employees {
		if emp.ID == id {
			return emp, true
		}
	}
	return Employee{}, false
}

func updateEmployee(id int, newName string) bool {
	for i, emp := range employees {
		if emp.ID == id {
			employees[i].Name = newName
			return true
		}
	}
	return false
}

func deleteEmployee(id int) bool {
	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	addEmployee(1, "Anita")
	addEmployee(2, "Ravi")
	fmt.Println("After adding:", employees)

	if emp, found := getEmployee(2); found {
		fmt.Println("Found:", emp)
	}

	updateEmployee(2, "Ravi Kumar")
	fmt.Println("After updating:", employees)

	deleteEmployee(1)
	fmt.Println("After deleting:", employees)
}
