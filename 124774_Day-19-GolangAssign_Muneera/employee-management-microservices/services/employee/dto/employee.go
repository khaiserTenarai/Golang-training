package dto

type CreateEmployeeRequest struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required,email"`
	Salary       float64 `json:"salary" binding:"required,gt=0"`
	DepartmentID uint    `json:"department_id" binding:"required,gt=0"`
}

type UpdateEmployeeRequest struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required,email"`
	Salary       float64 `json:"salary" binding:"required,gt=0"`
	DepartmentID uint    `json:"department_id" binding:"required,gt=0"`
}

type EmployeeResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Salary       float64 `json:"salary"`
	DepartmentID uint    `json:"department_id"`
}
