package main

import "fmt"

type Employee7 struct {
	ID     int
	Name   string
	Salary float64
}

func (employee *Employee7) incrSalary(amount float64) {
	employee.Salary = employee.Salary + amount
}
func main() {
	employee := Employee7{
		ID:     1,
		Name:   "Abirami",
		Salary: 80000,
	}
	fmt.Println("Salary Before Increment: ", employee.Salary)
	employee.incrSalary(10000)
	fmt.Println("Salary After Increment: ", employee.Salary)
}
