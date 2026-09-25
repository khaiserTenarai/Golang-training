// 10. Implement employee CRUD using a slice.

package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

var employees []Employee

func main() {
	// Create
	addEmployee(1, "Anita")
	addEmployee(2, "Ravi")
	addEmployee(3, "Priya")
	fmt.Println("After adding employees:")
	printEmployees()

	// Read
	fmt.Println()
	emp, found := getEmployee(2)
	if found {
		fmt.Println("Found employee:", emp.Name)
	}

	// Update
	updateEmployee(2, "Ravi Kumar")
	fmt.Println()
	fmt.Println("After updating employee 2:")
	printEmployees()

	// Delete
	deleteEmployee(1)
	fmt.Println()
	fmt.Println("After deleting employee 1:")
	printEmployees()
}

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

func updateEmployee(id int, newName string) {
	for i, emp := range employees {
		if emp.ID == id {
			employees[i].Name = newName
			return
		}
	}
}

func deleteEmployee(id int) {
	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return
		}
	}
}

func printEmployees() {
	for _, emp := range employees {
		fmt.Println(emp.ID, emp.Name)
	}
}
