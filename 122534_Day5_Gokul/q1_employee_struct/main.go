/*
Day 5, Q1. Create an Employee struct with at least 8 fields.
*/
package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Department string
	Designation string
	IsActive   bool
}

func main() {
	gokul := Employee{
		ID:          1,
		Name:        "Gokul",
		Email:       "gokul@example.com",
		Age:         27,
		Salary:      65000,
		Department:  "Data Engineering",
		Designation: "Data Engineer",
		IsActive:    true,
	}

	fmt.Printf("%+v\n", gokul)
}
