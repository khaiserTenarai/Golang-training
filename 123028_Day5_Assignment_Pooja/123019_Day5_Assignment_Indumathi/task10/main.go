package main

import (
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Employee struct {
	Person
	ContactInfo
	
	ID       int       `json:"id"`
	JobTitle string    `json:"job_title"`
	Salary   float64   `json:"salary"`
	IsActive bool      `json:"is_active"`
}

func main() {
	emp := Employee{
		Person: Person{
			FirstName: "Jane",
			LastName:  "Doe",
		},
		ContactInfo: ContactInfo{
			Email: "jane.doe@example.com",
			Phone: "555-0199",
		},
		ID:       101,
		JobTitle: "Software Engineer",
		Salary:   100000.0,
		IsActive: true,
	}

	fmt.Println("Full Name:", emp.FullName())
	fmt.Println("Email:", emp.Email)
	fmt.Println("Job Title:", emp.JobTitle)
}