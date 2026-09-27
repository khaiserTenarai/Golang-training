// Package model holds the core data types used across the app.
package model

// Employee represents one employee record.
type Employee struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Salary float64
}
