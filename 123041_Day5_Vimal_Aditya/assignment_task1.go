package main

import "fmt"

type Employee1 struct {
	ID         int
	Name       string
	Email      string
	Department string
	Designation string
	Age        int
	Salary     float64
	IsActive   bool
}

func main() {

	fmt.Println("\n**************************************************")
	fmt.Println("1. Create an Employee struct with at least 8 fields.")
	fmt.Println("****************************************************")

	emp := Employee1{
		ID:          101,
		Name:        "Vimal Aditya",
		Email:       "vimal@gmail.com",
		Department:  "Engineering",
		Designation: "Software Engineer",
		Age:         23,
		Salary:      500000,
		IsActive:    true,
	}

	fmt.Println("ID         :", emp.ID)
	fmt.Println("Name       :", emp.Name)
	fmt.Println("Email      :", emp.Email)
	fmt.Println("Department :", emp.Department)
	fmt.Println("Designation:", emp.Designation)
	fmt.Println("Age        :", emp.Age)
	fmt.Println("Salary     :", emp.Salary)
	fmt.Println("Is Active  :", emp.IsActive)
}