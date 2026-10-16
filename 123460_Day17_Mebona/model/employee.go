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

// EmployeeFilter describes HOW the employee list should be returned.
// It is shared by REST and GraphQL, so both behave the same way.
type EmployeeFilter struct {
	// filtering
	DepartmentID uint
	Search       string   // matches name or email
	MinSalary    *float64 // nil = not used
	MaxSalary    *float64 // nil = not used
	// sorting
	SortBy    string // id, name, email, salary, createdAt
	SortOrder string // asc or desc
	// pagination
	Page     int // starts at 1
	PageSize int // 1 to 100
}

// EmployeePage is one page of results plus the information needed to build "next / previous" buttons.
type EmployeePage struct {
	Items      []Employee `json:"items"`
	TotalCount int        `json:"totalCount"` // all matching employees (not just this page)
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	TotalPages int        `json:"totalPages"`
	HasNext    bool       `json:"hasNext"`
}
