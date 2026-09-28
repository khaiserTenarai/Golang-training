package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary = e.Salary + amount
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Swathi",
		Salary: 50000,
	}

	fmt.Println("Before salary increase:", employee.Salary)

	employee.IncreaseSalary(5000)

	fmt.Println("After salary increase:", employee.Salary)
}