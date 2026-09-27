// 1. Create functions for employee CRUD.

package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var store = make(map[int]Employee)

func createEmployee(emp Employee) {
	store[emp.ID] = emp
}

func getEmployee(id int) (Employee, bool) {
	emp, ok := store[id]
	return emp, ok
}

func updateEmployee(id int, name string, salary float64) bool {
	if _, ok := store[id]; !ok {
		return false
	}
	store[id] = Employee{ID: id, Name: name, Salary: salary}
	return true
}

func deleteEmployee(id int) bool {
	if _, ok := store[id]; !ok {
		return false
	}
	delete(store, id)
	return true
}

func main() {
	createEmployee(Employee{ID: 1, Name: "Alice", Salary: 50000})
	createEmployee(Employee{ID: 2, Name: "Bob", Salary: 60000})

	emp, ok := getEmployee(1)
	if ok {
		fmt.Println("Found:", emp)
	}

	updateEmployee(1, "Alice Smith", 55000)
	fmt.Println("Updated:", store[1])

	deleteEmployee(2)
	fmt.Println("Store:", store)
}
