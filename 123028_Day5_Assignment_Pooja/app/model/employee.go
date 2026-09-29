package model

import "time"

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Employee struct {
	Person
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	JobTitle string    `json:"job_title"`
	Salary   float64   `json:"salary"`
	HireDate time.Time `json:"hire_date"`
	IsActive bool      `json:"is_active"`
}