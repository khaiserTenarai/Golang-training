package main

import "fmt"

type Employee8 struct {
	ID     int
	Name   string
	Salary float64
}

func (employee Employee8) incrSalary(amount float64) {
	employee.Salary = employee.Salary + amount
	fmt.Println("Value of Salary inside method: ", employee.Salary)
}
func main() {
	employee := Employee8{
		ID:     1,
		Name:   "Abirami",
		Salary: 80000,
	}
	fmt.Println("Salary Before Increment: ", employee.Salary)
	employee.incrSalary(10000)
	fmt.Println("Salary After calling the incrSalary method: ", employee.Salary)
}
