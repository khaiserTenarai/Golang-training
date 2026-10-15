package model

import "time"

type Employee struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	Name         string      `gorm:"size:100;not null" json:"name"`
	Email        string      `gorm:"size:150;not null;uniqueIndex" json:"email"`
	Salary       float64     `gorm:"not null" json:"salary"`
	DepartmentID uint        `gorm:"not null;index" json:"departmentId"`
	Department   *Department `json:"department,omitempty"`
	CreatedAt    time.Time   `json:"createdAt"`
}

// EmployeeFilter is used by the "employees" list query (query arguments).
type EmployeeFilter struct {
	DepartmentID uint
	Search       string
	Limit        int
	Offset       int
}
