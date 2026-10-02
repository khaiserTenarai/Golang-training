package models

import "time"

type Employee struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	Salary       float64   `json:"salary"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Pincode      string    `json:"pincode"`
	DepartmentID *int      `json:"department_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EmployeeWithDepartment struct {
	Employee
	DepartmentName string `json:"department_name"`
	DepartmentCode string `json:"department_code"`
}