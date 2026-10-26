package dto

import "time"

type CreateEmployeeRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int64  `json:"department_id" binding:"required"`
}

type UpdateEmployeeRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DepartmentID int64  `json:"department_id"`
	Status       string `json:"status"`
}

type EmployeeResponse struct {
	ID           int64               `json:"id"`
	FirstName    string              `json:"first_name"`
	LastName     string              `json:"last_name"`
	Email        string              `json:"email"`
	DepartmentID int64               `json:"department_id"`
	Status       string              `json:"status"`
	Department   *DepartmentDTO      `json:"department,omitempty"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}

type DepartmentDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Headcount int    `json:"headcount_limit"`
}
