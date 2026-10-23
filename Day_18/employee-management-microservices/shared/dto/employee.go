package dto

// EmployeeDTO is the employee response shared through APIs.
type EmployeeDTO struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	DepartmentID uint    `json:"departmentId"`
	Salary       float64 `json:"salary"`
}

// CreateEmployeeRequest contains input for creating an employee.
type CreateEmployeeRequest struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	DepartmentID uint    `json:"departmentId" binding:"required"`
	Salary       float64 `json:"salary"`
}

// UpdateEmployeeRequest contains input for updating an employee.
type UpdateEmployeeRequest struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	DepartmentID uint    `json:"departmentId" binding:"required"`
	Salary       float64 `json:"salary"`
}
