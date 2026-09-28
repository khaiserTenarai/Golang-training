package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Phone      string
	Department string
	Address    string
}

func main() {
	employee := Employee{
		ID:         1,
		Name:       "Abirami",
		Email:      "abi@gmail.com",
		Age:        22,
		Salary:     80000,
		Phone:      "9876543210",
		Department: "IT",
		Address:    "Banglore",
	}
	fmt.Println(employee)
}
