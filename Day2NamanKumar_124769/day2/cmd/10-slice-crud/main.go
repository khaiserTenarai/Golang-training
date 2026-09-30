// Task 10: Employee CRUD (Create, Read, Update, Delete) using a SLICE.
// A slice is a dynamic, growable view over an array.
package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var ErrNotFound = errors.New("employee not found")

// employees is the slice that stores all records.
var employees []Employee

// Create adds a new employee to the end of the slice.
func Create(e Employee) error {
	if _, idx := findIndex(e.ID); idx != -1 {
		return fmt.Errorf("ID %d already exists", e.ID)
	}
	employees = append(employees, e)
	return nil
}

// Read returns the employee with the given ID.
func Read(id int) (Employee, error) {
	e, idx := findIndex(id)
	if idx == -1 {
		return Employee{}, ErrNotFound
	}
	return e, nil
}

// Update changes the salary and department of an existing employee.
func Update(id int, dept string, salary float64) error {
	_, idx := findIndex(id)
	if idx == -1 {
		return ErrNotFound
	}
	// Modify through the index. Changing the range copy would NOT update the slice.
	employees[idx].Department = dept
	employees[idx].Salary = salary
	return nil
}

// Delete removes an employee while keeping the order of the rest.
func Delete(id int) error {
	_, idx := findIndex(id)
	if idx == -1 {
		return ErrNotFound
	}
	employees = append(employees[:idx], employees[idx+1:]...)
	return nil
}

// findIndex does a linear search: O(n). (Task 11 shows a map, which is O(1).)
func findIndex(id int) (Employee, int) {
	for i, e := range employees {
		if e.ID == id {
			return e, i
		}
	}
	return Employee{}, -1
}

func list(title string) {
	fmt.Printf("\n--- %s (len=%d cap=%d) ---\n", title, len(employees), cap(employees))
	for i, e := range employees {
		fmt.Printf("[%d] %d  %-12s %-12s ₹%.0f\n", i, e.ID, e.Name, e.Department, e.Salary)
	}
}

func main() {
	fmt.Println("=== CREATE ===")
	for _, e := range []Employee{
		{101, "Anita Rao", "Engineering", 90000},
		{102, "Rahul Mehta", "HR", 45000},
		{103, "Priya Nair", "Finance", 60000},
		{104, "Arjun Singh", "Engineering", 70000},
	} {
		if err := Create(e); err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("added %-12s len=%d cap=%d\n", e.Name, len(employees), cap(employees))
	}
	fmt.Println("(cap grows 1 -> 2 -> 4: Go doubles the underlying array when it's full)")
	fmt.Println("duplicate:", Create(Employee{ID: 101, Name: "Duplicate"}))
	list("All employees")

	fmt.Println("\n=== READ ===")
	if e, err := Read(103); err == nil {
		fmt.Printf("found: %+v\n", e)
	}
	_, err := Read(999)
	fmt.Println("read 999:", err)

	fmt.Println("\n=== UPDATE ===")
	fmt.Println("update 102:", Update(102, "Operations", 52000))
	fmt.Println("update 999:", Update(999, "X", 1))
	list("After update")

	fmt.Println("\n=== DELETE ===")
	fmt.Println("delete 101:", Delete(101))
	fmt.Println("delete 999:", Delete(999))
	list("After delete")

	fmt.Println("\n=== Slice tricks ===")
	fmt.Println("first 2 :", names(employees[:2]))
	copied := make([]Employee, len(employees))
	n := copy(copied, employees)
	copied[0].Name = "Changed"
	fmt.Printf("copy()  : copied %d; original[0] is still %q\n", n, employees[0].Name)
}

func names(es []Employee) []string {
	out := make([]string, 0, len(es))
	for _, e := range es {
		out = append(out, e.Name)
	}
	return out
}
