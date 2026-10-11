// Package model contains application domain models.
package model

// Employee represents an employee stored in PostgreSQL.
type Employee struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// EmployeeFilter encapsulates pagination, filtering, and sorting parameters.
type EmployeeFilter struct {
	Department string
	SortBy     string
	Order      string
	Page       int
	Limit      int
}