package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func main() {
	employee := Employee{
		ID:   101,
		Name: "Abirami",
	}

	fmt.Printf("Employee ID: %d\n", employee.ID)
	fmt.Printf("Employee Name: %s\n", employee.Name)
}
