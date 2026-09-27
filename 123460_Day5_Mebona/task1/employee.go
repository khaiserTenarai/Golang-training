package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Phone      string
	Position   string
	Experience int
}

func main() {
	employee := Employee{
		ID:         101,
		Name:       "John",
		Email:      "john@gmail.com",
		Age:        28,
		Salary:     50000,
		Phone:      "9876543210",
		Position:   "Developer",
		Experience: 5,
	}

	fmt.Println(employee)
}