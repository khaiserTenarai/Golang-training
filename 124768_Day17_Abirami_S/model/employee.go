package model

import "gorm.io/gorm"

type Department struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Employees []Employee `json:"employees,omitempty"`
}

type Employee struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Email        string         `json:"email" gorm:"unique"`
	Salary       float64        `json:"salary"`
	DepartmentID uint           `json:"departmentId"`
	Department   *Department    `json:"department,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type EmployeeFilter struct {
	Name string
}

type SortOption struct {
	Field string
	Order string
}
