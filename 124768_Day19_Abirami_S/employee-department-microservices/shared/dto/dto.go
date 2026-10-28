package dto

type CreateEmployeeRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int    `json:"department_id" binding:"required,gt=0"`
}

type UpdateEmployeeRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int    `json:"department_id" binding:"required,gt=0"`
}

type EmployeeResponse struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Email        string              `json:"email"`
	DepartmentID int                 `json:"department_id"`
	Department   *DepartmentResponse `json:"department,omitempty"`
}

type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type DepartmentResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Error         string `json:"error"`
	CorrelationID string `json:"correlation_id,omitempty"`
}
