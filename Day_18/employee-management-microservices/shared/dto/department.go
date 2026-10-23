package dto

// DepartmentDTO is the department response shared through APIs.
type DepartmentDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateDepartmentRequest contains input for creating a department.
type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// UpdateDepartmentRequest contains input for updating a department.
type UpdateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
