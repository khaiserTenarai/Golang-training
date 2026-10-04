package model

import "time"

type Attendance struct {
	ID         int
	EmployeeID int
	CheckIn    time.Time
	CheckOut   *time.Time
}
