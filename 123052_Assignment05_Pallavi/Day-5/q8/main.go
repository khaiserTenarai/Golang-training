package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Value receiver
func (e Employee) changeName(name string) {

	// This changes only the copy
	e.Name = name

	fmt.Println("Inside Method:", e.Name)
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Pallavi",
		Salary: 29000,
	}

	fmt.Println("Before Method:", employee.Name)

	// Calling value receiver
	employee.changeName("Pallavi K")

	// Original value is not changed
	fmt.Println("After Method:", employee.Name)
}