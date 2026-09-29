package main

import "fmt"

type Employee struct {
	ID           int
	Name         string
	Email        string
	Age          int
	Department   string
	Position     string
	Salary       float64
	MobileNumber string
}

func main() {
	employee := Employee{
		ID:           101,
		Name:         "Shashank Raj",
		Email:        "shank.com",
		Age:          23,
		Department:   "Engineering",
		Position:     "SDE",
		Salary:       99990,
		MobileNumber: "9949494940",
	}

	fmt.Println(employee)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Department:", employee.Department)
}
