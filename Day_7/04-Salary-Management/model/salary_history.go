package model
import "time"

type SalaryHistory struct {
	ID         int
	EmployeeID int
	OldSalary  float64
	NewSalary  float64
	ChangedAt  time.Time
}
