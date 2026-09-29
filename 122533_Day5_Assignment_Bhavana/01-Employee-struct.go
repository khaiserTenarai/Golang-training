// 1. Create an Employee struct with at least 8 fields.

package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Department string
	Phone      string
	IsActive   bool
}

func main() {
	emp := Employee{
		ID:         1,
		Name:       "Anita",
		Email:      "anita@example.com",
		Age:        28,
		Salary:     45000,
		Department: "Engineering",
		Phone:      "9876543210",
		IsActive:   true,
	}

	fmt.Println(emp)
	fmt.Printf("%+v\n", emp)
}
