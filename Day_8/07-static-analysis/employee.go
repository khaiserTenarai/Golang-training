package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func (e Employee) Display() {
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Department: %s\n", e.Department)
	fmt.Printf("Salary: %.2f\n", e.Salary)
}

func main() {
	employee := Employee{
		ID:         101,
		Name:       "Alice",
		Department: "Engineering",
		Salary:     75000,
	}

	employee.Display()
}
