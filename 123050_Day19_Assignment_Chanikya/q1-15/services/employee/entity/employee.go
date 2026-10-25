package entity

import "time"

type Employee struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Email        string    `gorm:"size:150;not null;uniqueIndex" json:"email"`
	Salary       float64   `gorm:"not null" json:"salary"`
	DepartmentID uint      `gorm:"not null;index" json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Employee) TableName() string {
	return "employees"
}
