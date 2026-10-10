// Package model defines application models.
package model

// Employee represents an employee.
type Employee struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}
