// Tasks 5, 6: DTOs — Data Transfer Objects for API communication.
// DTOs are separate from entities. They control what the API exposes.

package dto

type EmployeeResponse struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Email        string              `json:"email"`
	Age          int                 `json:"age"`
	DepartmentID int                 `json:"department_id"`
	Salary       float64             `json:"salary"`
	IsActive     bool                `json:"is_active"`
	CreatedAt    string              `json:"created_at"`
	Department   *DepartmentResponse `json:"department,omitempty"`
}

type CreateEmployeeRequest struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	DepartmentID int     `json:"department_id"`
	Salary       float64 `json:"salary"`
}

type UpdateEmployeeRequest struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	DepartmentID int     `json:"department_id"`
	Salary       float64 `json:"salary"`
	IsActive     bool    `json:"is_active"`
}

type DepartmentResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Budget   float64 `json:"budget"`
}

type CreateDepartmentRequest struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Budget   float64 `json:"budget"`
}

type APIResponse struct {
	Success       bool        `json:"success"`
	Message       string      `json:"message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	CorrelationID string      `json:"correlation_id,omitempty"`
}

type HealthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

type ErrorResponse struct {
	Success       bool   `json:"success"`
	Error         string `json:"error"`
	Code          int    `json:"code"`
	CorrelationID string `json:"correlation_id,omitempty"`
}
