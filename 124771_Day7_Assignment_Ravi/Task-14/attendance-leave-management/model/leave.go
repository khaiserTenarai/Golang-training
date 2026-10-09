package model

import "time"

type Leave struct {
    ID         int
    EmployeeID int
    LeaveDate  time.Time
    Reason     string
    Status     string
}
