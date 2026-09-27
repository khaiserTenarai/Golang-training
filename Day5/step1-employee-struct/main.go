package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Salary   float64
	City     string
	Pincode  string
	DeptName string
}

func main() {
	e := Employee{
		ID:       1,
		Name:     "Ray",
		Email:    "ray@example.com",
		Age:      26,
		Salary:   60000,
		City:     "Bengaluru",
		Pincode:  "560001",
		DeptName: "Engineering",
	}

	fmt.Println("Employee struct with 8 fields:")
	fmt.Printf("%+v\n", e)
}
