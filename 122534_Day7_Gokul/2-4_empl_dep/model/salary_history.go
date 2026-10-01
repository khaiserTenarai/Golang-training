package model

import "time"

// SalaryHistory mirrors one row in the "salary_history" table - a
// single record of an employee's salary changing from one value to
// another.
type SalaryHistory struct {
	ID         int
	EmployeeID int
	OldSalary  float64
	NewSalary  float64
	ChangedAt  time.Time
}
