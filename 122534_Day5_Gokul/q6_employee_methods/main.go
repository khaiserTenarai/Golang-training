
package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) IsEligibleForBonus() bool {
	return e.Salary >= 50000
}

func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}

func main() {
	gokul := Employee{Name: "Gokul", Salary: 65000}

	fmt.Println("Is Eligible for bonus", gokul.IsEligibleForBonus())
	fmt.Println("Annual salary :", gokul.AnnualSalary())
}
