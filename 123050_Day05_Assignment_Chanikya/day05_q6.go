package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Method to display employee details
func (e Employee) Display() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
}

// Method to increase salary
func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary += amount
}

func main() {
	employee := Employee{
		ID:     101,
		Name:   "Ram",
		Salary: 50000,
	}

	employee.Display()

	employee.IncreaseSalary(5000)

	fmt.Println("New Salary:", employee.Salary)
}
