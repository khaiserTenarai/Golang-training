package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Describe() string {
	return fmt.Sprintf("%s earns %.2f per month", e.Name, e.Salary)
}

func main() {
	gokul := Employee{Name: "Gokul", Salary: 65000}
	fmt.Println(gokul.Describe())
}
