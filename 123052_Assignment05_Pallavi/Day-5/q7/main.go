package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Pointer receiver
func (e *Employee) changeName(name string) {

	// Modify original employee
	e.Name = name
}

// Pointer receiver to increase salary
func (e *Employee) increaseSalary(amount float64) {

	e.Salary = e.Salary + amount
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Pallavi",
		Salary: 29000,
	}

	fmt.Println("Before Change")
	fmt.Println("Name:", employee.Name)
	fmt.Println("Salary:", employee.Salary)

	// Calling pointer receiver
	employee.changeName("Pallavi K")
	employee.increaseSalary(5000)

	fmt.Println("\nAfter Change")
	fmt.Println("Name:", employee.Name)
	fmt.Println("Salary:", employee.Salary)
}