package models

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Role       string  `json:"role" binding:"required"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"required,gt=0"`
}

type UpdateEmployee struct {
	Name       *string  `json:"name,omitempty"`
	Role       *string  `json:"role,omitempty"`
	Department *string  `json:"department,omitempty"`
	Salary     *float64 `json:"salary,omitempty"`
}
