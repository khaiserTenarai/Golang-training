package model

import "time"

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"size:120;not null" json:"name"`
	Email        string     `gorm:"size:180;uniqueIndex;not null" json:"email"`
	Salary       float64    `gorm:"not null" json:"salary"`
	DepartmentID uint       `gorm:"not null;index" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
