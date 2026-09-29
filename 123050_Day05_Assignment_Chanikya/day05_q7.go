package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

// Pointer receiver
func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary += amount
}

func main() {
	employee := Employee{
		Name:   "Ram",
		Salary: 50000,
	}

	fmt.Println("Before:", employee.Salary)

	employee.IncreaseSalary(5000)

	fmt.Println("After:", employee.Salary)
}
