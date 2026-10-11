// Package model contains domain models.
package model

// Employee represents an employee model in PostgreSQL.
type Employee struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// PartialEmployee represents a payload for partial updates (PATCH).
type PartialEmployee struct {
	Name       *string  `json:"name"`
	Email      *string  `json:"email"`
	Department *string  `json:"department"`
	Salary     *float64 `json:"salary"`
}