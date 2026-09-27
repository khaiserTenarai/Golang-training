package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e Employee) GetAnnualSalary() float64 {
	return e.Salary * 12
}

func (e Employee) Display() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Piyush",
		Salary: 50000,
	}

	employee.Display()

	fmt.Println("Annual Salary:", employee.GetAnnualSalary())
}