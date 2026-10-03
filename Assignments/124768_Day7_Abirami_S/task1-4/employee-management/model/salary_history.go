package model

import "time"

// task4
type SalaryHistory struct {
	ID         int
	EmployeeID int
	OldSalary  float64
	NewSalary  float64
	ChangedAt  time.Time
}
