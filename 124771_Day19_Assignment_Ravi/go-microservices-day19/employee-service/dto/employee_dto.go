package dto

// CreateEmployeeRequest is input validation, separate from the GORM entity.
type CreateEmployeeRequest struct {
	Name         string `json:"name" binding:"required,min=2,max=120"`
	Email        string `json:"email" binding:"required,email,max=200"`
	DepartmentID uint   `json:"department_id" binding:"required,gt=0"`
}
type UpdateEmployeeRequest struct {
	Name         string `json:"name" binding:"required,min=2,max=120"`
	Email        string `json:"email" binding:"required,email,max=200"`
	DepartmentID uint   `json:"department_id" binding:"required,gt=0"`
}

// DepartmentDTO represents data returned by the other service.
type DepartmentDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// EmployeeResponse composes local employee data and remote department data.
type EmployeeResponse struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	DepartmentID uint           `json:"department_id"`
	Department   *DepartmentDTO `json:"department,omitempty"`
}
