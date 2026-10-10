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

// EmployeeListResponse represents a paginated employee response.
type EmployeeListResponse struct {
	Data       []Employee `json:"data"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	Total      int        `json:"total"`
	TotalPages int        `json:"totalPages"`
}
