package model

import "time"

// json tags match the GraphQL field names, so graphql-go's default resolver can read them.
type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"size:100;not null" json:"name"`
	Email        string     `gorm:"size:150;not null;uniqueIndex" json:"email"`
	Salary       float64    `gorm:"not null" json:"salary"`
	Active       bool       `gorm:"not null" json:"active"`
	JoinedAt     time.Time  `gorm:"not null" json:"joinedAt"`
	DepartmentID uint       `gorm:"not null;index" json:"departmentId"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"` // belongs-to
}

// CreateEmployeeInput is used by REST (JSON body) and GraphQL (input type).
type CreateEmployeeInput struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Salary       float64 `json:"salary"`
	DepartmentID uint    `json:"departmentId"`
	Active       *bool   `json:"active"` // nil = default (true)
}

// UpdateEmployeeInput: nil = "field not sent, keep the old value" (partial update).
type UpdateEmployeeInput struct {
	Name         *string  `json:"name"`
	Email        *string  `json:"email"`
	Salary       *float64 `json:"salary"`
	DepartmentID *uint    `json:"departmentId"`
	Active       *bool    `json:"active"`
}

// EmployeeFilter holds the arguments of the `employees` query.
type EmployeeFilter struct {
	DepartmentID *uint
	Search       string
	JoinedAfter  *time.Time
	Limit        int // 0 = no limit
	Offset       int
}
