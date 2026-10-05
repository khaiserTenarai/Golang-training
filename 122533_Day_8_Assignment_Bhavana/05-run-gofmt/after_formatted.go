package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	emp := Employee{Name: "Ram", Salary: 45000}
	if emp.Salary > 40000 {
		fmt.Println("High earner:", emp.Name)
	}
}
