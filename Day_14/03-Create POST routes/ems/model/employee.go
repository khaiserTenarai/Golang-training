// Package model defines employee data.
package model

// Employee represents an employee record.
type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
