package model

import "time"

// Employee is a persistence entity. The department's name is not stored here;
// it is fetched from Department Service over REST (microservice boundary).
type Employee struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:120;not null" json:"name"`
	Email        string    `gorm:"size:200;not null;uniqueIndex" json:"email"`
	DepartmentID uint      `gorm:"not null;index" json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
