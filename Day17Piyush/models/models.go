package models

import "time"

type Employee struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	DepartmentID int       `json:"department_id"`
	Salary       float64   `json:"salary"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type PaginationInput struct {
	Page     int
	PageSize int
}

type EmployeeFilter struct {
	Name       string
	Department string
	MinSalary  float64
	MaxSalary  float64
	SortBy     string
	SortOrder  string
}
