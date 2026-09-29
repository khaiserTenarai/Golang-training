package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Age        int
	Email      string
	Department string
	Salary     float64
	Phone      string
	Address    string
	IsActive   bool
}

func main() {
	employee := Employee{
		ID:         101,
		Name:       "ram",
		Age:        24,
		Email:      "ram@example.com",
		Department: "Software Engineer",
		Salary:     50000,
		Phone:      "9876543210",
		Address:    "Bangalore",
		IsActive:   true,
	}

	fmt.Println(employee)
}
