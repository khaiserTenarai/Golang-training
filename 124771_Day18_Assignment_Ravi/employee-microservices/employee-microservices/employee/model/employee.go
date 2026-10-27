package model

// Employee is the Employee Service's local representation.
// Task 3: bounded context. Task 4: Employee Service.
type Employee struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	DepartmentID int    `json:"departmentId"`
}
