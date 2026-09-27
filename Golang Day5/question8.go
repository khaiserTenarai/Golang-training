package main

import "fmt"

type Employee struct {
	FirstName     string
	LastName      string
	MonthlySalary int
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) AnnualSalary() int {
	return e.MonthlySalary * 12
}

func (e Employee) AttemptSalaryIncrease() {
	e.MonthlySalary += 1000
}

func main() {
	emp := Employee{
		FirstName:     "Lakshmi",
		LastName:      "Shibu",
		MonthlySalary: 5000,
	}

	fmt.Println(emp.FullName())
	fmt.Println(emp.AnnualSalary())

	emp.AttemptSalaryIncrease()
	fmt.Println(emp.MonthlySalary)
}