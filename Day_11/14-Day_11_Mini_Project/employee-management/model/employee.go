// Package model contains application domain models.
package model

// Employee represents an employee stored in PostgreSQL.
type Employee struct {
	// ID is the unique employee identifier.
	ID int64 `json:"id"`
	// Name contains the employee name.
	Name string `json:"name"`
	// Email contains the employee email address.
	Email string `json:"email"`
	// Department contains the employee department.
	Department string `json:"department"`
	// Salary contains the employee salary.
	Salary float64 `json:"salary"`
}
