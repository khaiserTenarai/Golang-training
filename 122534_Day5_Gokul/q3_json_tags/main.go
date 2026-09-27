package main

import "fmt"

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
	SSN    string  `json:"-"`
}

func main() {
	gokul := Employee{
		ID:     1,
		Name:   "Gokul",
		Email:  "gokul@example.com",
		Salary: 65000,
		SSN:    "hidden-value",
	}

	fmt.Printf("%+v\n", gokul)
}
