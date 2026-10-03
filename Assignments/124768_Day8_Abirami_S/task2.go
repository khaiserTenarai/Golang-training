package main

import (
	"fmt"
	"log"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {
	employee := Employee{
		ID:     101,
		Name:   "Tom",
		Salary: 70000,
	}
	log.Println("Employee created successfully")

	fmt.Println("Employee ID: ", employee.ID)
	fmt.Println("Employee Name: ", employee.Name)
	fmt.Println("Basic Salary: ", employee.Salary)

	log.Println("Application completed")
}
