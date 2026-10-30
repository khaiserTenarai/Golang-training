package entity

import "time"

type EmployeeStatus string

const (
	StatusPending  EmployeeStatus = "PENDING"
	StatusActive   EmployeeStatus = "ACTIVE"
	StatusRejected EmployeeStatus = "REJECTED"
)

type Employee struct {
	ID           int64          `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email"`
	DepartmentID int64          `json:"department_id"`
	Status       EmployeeStatus `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
