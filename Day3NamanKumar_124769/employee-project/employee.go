package main

import "fmt"

// Employee models a single employee record in the company directory.
type Employee struct {
	ID     int
	Name   string
	Role   string
	Salary float64
}

// Employees is an in-memory store of employee records.
var Employees []Employee

// AddEmployee appends a new employee to the store.
func AddEmployee(id int, name, role string, salary float64) {
	Employees = append(Employees, Employee{ID: id, Name: name, Role: role, Salary: salary})
}

// ListEmployees prints all employees to stdout.
func ListEmployees() {
	for _, e := range Employees {
		fmt.Printf("Employee #%d: %s | %s | $%.2f\n", e.ID, e.Name, e.Role, e.Salary)
	}
}

// RemoveEmployee removes an employee by ID.
func RemoveEmployee(id int) bool {
	for i, e := range Employees {
		if e.ID == id {
			Employees = append(Employees[:i], Employees[i+1:]...)
			return true
		}
	}
	return false
}
