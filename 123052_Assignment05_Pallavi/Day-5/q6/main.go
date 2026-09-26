package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Method to display employee
func (e Employee) displayEmployee() {

	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Salary:", e.Salary)
}

// Method to increase salary
func (e Employee) calculateBonus() float64 {

	return e.Salary * 0.10
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Pallavi",
		Salary: 29000,
	}

	// Calling employee method
	employee.displayEmployee()

	bonus := employee.calculateBonus()

	fmt.Println("Bonus:", bonus)
}