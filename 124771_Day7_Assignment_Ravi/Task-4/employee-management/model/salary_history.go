package model

import "time"

type SalaryHistory struct {
	ID         int64
	EmployeeID int64
	OldSalary  float64
	NewSalary  float64
	ChangedAt  time.Time
}
