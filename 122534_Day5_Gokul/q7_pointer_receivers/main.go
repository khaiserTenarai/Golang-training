package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e *Employee) Promote(raise float64) {
	e.Salary += raise
}

func main() {
	gokul := Employee{Name: "Gokul", Salary: 65000}

	fmt.Println("Before promotion:", gokul.Salary)
	gokul.Promote(8000)
	fmt.Println("After promotion :", gokul.Salary)
}
