// Task 6: Entities — internal domain models (separate from DTOs).
// Entities represent core data as stored internally.

package entity

import "time"

type Employee struct {
	ID           int
	Name         string
	Email        string
	Age          int
	DepartmentID int
	Salary       float64
	IsActive     bool
	CreatedAt    time.Time
}

type Department struct {
	ID       int
	Name     string
	Location string
	Budget   float64
}
