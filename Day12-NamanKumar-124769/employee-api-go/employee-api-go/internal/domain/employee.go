package domain

import "time"

type Employee struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateEmployeeInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

type UpdateEmployeeInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

type EmployeeQuery struct {
	Page       int
	Limit      int
	Department string
	Search     string
	SortBy     string
	Order      string
}

type EmployeePage struct {
	Data       []Employee `json:"data"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
}
