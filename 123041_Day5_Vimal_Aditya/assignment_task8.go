package main

import "fmt"

type Employee8 struct {
	Name   string
	Salary float64
}

func (e Employee8) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
	fmt.Println(e.Salary)
}

func main() {

	fmt.Println("\n********************************")
	fmt.Println("8. Demonstrate value receivers.")
	fmt.Println("*******************************")

	emp := Employee8{
		Name:   "Vimal",
		Salary: 500000,
	}

	fmt.Println("Initial Salary    :", emp.Salary)

	emp.UpdateSalary(600000)

	fmt.Println("After Value Recv  :", emp.Salary)

}