package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

// Value receiver
func (e Employee) Display() {
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
}

func main() {
	employee := Employee{
		Name:   "Ram",
		Salary: 50000,
	}

	employee.Display()
}
