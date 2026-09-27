// 1. Create an Employee struct with at least 8 fields.

package main

import "fmt"

type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	Age        int
	Department string
	Salary     float64
	IsActive   bool
}

func main() {
	emp := Employee{
		ID:         101,
		FirstName:  "Sachin",
		LastName:   "M S",
		Email:      "sachin.ms@example.com",
		Age:        30,
		Department: "Engineering",
		Salary:     35000.50,
		IsActive:   true,
	}

	fmt.Println("Employee ID:", emp.ID)
	fmt.Println("Full Name:", emp.FirstName, emp.LastName)
	fmt.Println("Department:", emp.Department)
	fmt.Println("Active:", emp.IsActive)
}
