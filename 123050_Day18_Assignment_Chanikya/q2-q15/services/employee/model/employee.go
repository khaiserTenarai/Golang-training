package model

import "time"

// Employee represents an employee record.
type Employee struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Email        string    `gorm:"size:150;not null;uniqueIndex" json:"email"`
	DepartmentID uint      `gorm:"not null;index" json:"departmentId"`
	Salary       float64   `gorm:"not null" json:"salary"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
