package models

import "time"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	CreatedAt  time.Time
}

type SearchParams struct {
	Name       string
	Department string
	MinSalary  float64
	MaxSalary  float64
	SortBy     string
	SortOrder  string
	Page       int
	PageSize   int
}
