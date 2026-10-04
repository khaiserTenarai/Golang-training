package main

import "fmt"

type Employee7 struct {
	Name   string
	Salary float64
}

// Pointer receiver
func (e *Employee7) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

func main() {

	fmt.Println("\n*********************************")
	fmt.Println("7. Demonstrate pointer receivers.")
	fmt.Println("*********************************")

	emp := Employee7{
		Name:   "Sasi",
		Salary: 500000,
	}

	fmt.Println("Initial Salary    :", emp.Salary)

	emp.UpdateSalary(600000)

	fmt.Println("After Pointer Recv:", emp.Salary)
}