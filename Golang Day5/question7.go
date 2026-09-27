package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func (e Employee) FailedRaise() {
	e.Salary += 5000
}

func (e *Employee) SuccessfulRaise() {
	e.Salary += 5000
}

func main() {
	emp := Employee{
		Name:   "Lakshmi Shibu",
		Salary: 50000,
	}

	fmt.Println(emp.Name)
	fmt.Println("Starting Salary:", emp.Salary)

	emp.FailedRaise()
	fmt.Println("Salary after FailedRaise:", emp.Salary)

	emp.SuccessfulRaise()
	fmt.Println("Salary after SuccessfulRaise:", emp.Salary)
}