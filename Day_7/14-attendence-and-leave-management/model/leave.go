package model
import "time"

type Leave struct {
	ID         int
	EmployeeID int
	FromDate   time.Time
	ToDate     time.Time
	Reason     string
	Status     string
}

type LeaveReport struct {
	ID           int
	EmployeeID   int
	EmployeeName string
	FromDate     time.Time
	ToDate       time.Time
	Reason       string
	Status       string
}
