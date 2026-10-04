package model

import "time"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Department string
}

type Attendance struct {
	ID             int
	EmployeeID     int
	EmployeeName   string
	AttendanceDate time.Time
	CheckIn        *time.Time
	CheckOut       *time.Time
}

type Leave struct {
	ID           int
	EmployeeID   int
	EmployeeName string
	LeaveType    string
	StartDate    time.Time
	EndDate      time.Time
	Reason       string
	Status       string
	AppliedAt    time.Time
}
