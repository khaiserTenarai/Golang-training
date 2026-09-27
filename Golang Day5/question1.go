package main

import "fmt"

type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	JobTitle   string
	Department string
	Salary     float64
	IsActive   bool
}

func (e Employee) DisplayInfo() {
	fmt.Println("--- Employee Profile ---")
	fmt.Printf("Name: %s %s\n", e.FirstName, e.LastName)
	fmt.Printf("Employee ID: %d\n", e.ID)
	fmt.Printf("Role: %s - %s\n", e.JobTitle, e.Department)
	fmt.Printf("Contact: %s\n", e.Email)
	fmt.Printf("Salary: $%.2f\n", e.Salary)
	
	
	if e.IsActive {
		fmt.Println("Status: Currently Active")
	} else {
		fmt.Println("Status: Inactive/Former Employee")
	}
	
}

func main() {
	
	emp1 := Employee{
		ID:         1042,
		FirstName:  "Maya",
		LastName:   "Patel",
		Email:      "maya.p@example.com",
		JobTitle:   "Data Analyst",
		Department: "Analytics",
		Salary:     72500.00,
		IsActive:   true,
	}

	
	fmt.Println("Sending welcome email to:", emp1.Email)
	fmt.Println()

	
	emp1.DisplayInfo()

	
	emp2 := Employee{
		ID:         1043,
		FirstName:  "Lakshmi",
		LastName:   "Shibu",
		Email:      "lakshmi.shibu@gmail.com",
		JobTitle:   "Project Manager",
		Department: "Operations",
		Salary:     90000.00,
		IsActive:   false,
	}
	
	
	emp2.IsActive = true 
	
	emp2.DisplayInfo()
}