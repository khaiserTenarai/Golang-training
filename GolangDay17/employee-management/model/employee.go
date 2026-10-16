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

type EmployeeFilter struct {
	DepartmentID uint
	Search       string
	MinSalary    *float64
	MaxSalary    *float64
	SortBy       string
	SortOrder    string
	Page         int
	PageSize     int
}

type EmployeePage struct {
	Items      []Employee `json:"items"`
	TotalCount int        `json:"totalCount"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	TotalPages int        `json:"totalPages"`
	HasNext    bool       `json:"hasNext"`
}