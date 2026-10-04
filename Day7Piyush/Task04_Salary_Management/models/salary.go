package models

import "time"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	CreatedAt  time.Time
}

type SalaryHistory struct {
	ID           int
	EmployeeID   int
	EmployeeName string
	OldSalary    float64
	NewSalary    float64
	ChangeReason string
	ChangedAt    time.Time
}
