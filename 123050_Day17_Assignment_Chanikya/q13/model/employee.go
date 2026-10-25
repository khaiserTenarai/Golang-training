package model

// Employee is the database model shared by REST and GraphQL.
type Employee struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `gorm:"size:100;not null;index" json:"name"`
	Email      string  `gorm:"size:150;not null;uniqueIndex" json:"email"`
	Department string  `gorm:"size:100;not null;index" json:"department"`
	Salary     float64 `gorm:"not null" json:"salary"`
}

// EmployeeListResponse describes a paginated list returned by both APIs.
type EmployeeListResponse struct {
	Data       []Employee `json:"data"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	Total      int        `json:"total"`
	TotalPages int        `json:"totalPages"`
}

// EmployeeFilter contains the allowed employee filtering fields.
type EmployeeFilter struct {
	Name       string
	Department string
}

// EmployeeSort contains an allowlisted column and sort direction.
type EmployeeSort struct {
	Field     string
	Direction string
}
