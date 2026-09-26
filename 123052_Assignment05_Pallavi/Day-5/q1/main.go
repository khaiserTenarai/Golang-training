package main

import "fmt"

// Employee represents an employee
type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Department string
	Salary     float64
	Phone      string
	City       string
}

func main() {

	// Creating an Employee object
	employee := Employee{
		ID:         101,
		Name:       "Pallavi",
		Email:      "pallavi@gmail.com",
		Age:        22,
		Department: "IT",
		Salary:     29000,
		Phone:      "9876543210",
		City:       "Bangalore",
	}

	// Display employee details
	fmt.Println("Employee Details")
	fmt.Println("----------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Phone:", employee.Phone)
	fmt.Println("City:", employee.City)
}