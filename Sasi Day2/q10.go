package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func main() {
	employees := []Employee{}

	// CREATE
	employees = append(employees, Employee{1, "Sasikumar"})
	employees = append(employees, Employee{2, "Vijay"})

	// READ
	fmt.Println("Employees:")
	for _, employee := range employees {
		fmt.Println(employee.ID, employee.Name)
	}

	// UPDATE
	employees[0].Name = "Sasikumar Veeranan"

	// DELETE
	employees = append(employees[:1], employees[2:]...)

	fmt.Println("After Update and Delete:")

	for _, employee := range employees {
		fmt.Println(employee.ID, employee.Name)
	}
}