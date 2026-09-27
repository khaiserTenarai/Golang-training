package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// Value receiver
func (e Employee) Display() {
	fmt.Println("Employee Name  :", e.Name)
	fmt.Println("Employee Salary:", e.Salary)
}

func main() {

	employee := Employee{
		ID:       101,
		Name:     "Muneera",
		Position: "Software Developer",
		Salary:   20000,
	}

	// Calling value receiver method
	employee.Display()
}
