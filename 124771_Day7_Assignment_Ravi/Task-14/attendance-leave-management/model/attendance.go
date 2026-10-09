package model

import "time"

type Attendance struct {
    ID             int
    EmployeeID     int
    AttendanceDate time.Time
    CheckIn        *time.Time
    CheckOut       *time.Time
}
