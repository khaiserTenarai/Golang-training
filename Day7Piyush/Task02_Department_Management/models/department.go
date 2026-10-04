package models

import "time"

type Department struct {
	ID        int
	Name      string
	Location  string
	CreatedAt time.Time
}

type Employee struct {
	ID             int
	Name           string
	Email          string
	DepartmentID   *int
	DepartmentName string
	Salary         float64
	CreatedAt      time.Time
}
