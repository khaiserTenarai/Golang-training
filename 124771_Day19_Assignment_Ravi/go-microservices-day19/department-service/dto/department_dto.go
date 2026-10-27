package dto

// CreateDepartmentRequest is the public input contract (not a database entity).
type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=120"`
	Description string `json:"description" binding:"max=500"`
}

// UpdateDepartmentRequest uses the same validation for updates.
type UpdateDepartmentRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=120"`
	Description string `json:"description" binding:"max=500"`
}

// DepartmentResponse is the data exposed to API clients.
type DepartmentResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
