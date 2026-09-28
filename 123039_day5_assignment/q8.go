package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e Employee) Display() {
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Salary:", e.Salary)
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Swathi",
		Salary: 50000,
	}

	employee.Display()
}