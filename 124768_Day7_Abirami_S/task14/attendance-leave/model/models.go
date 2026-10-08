package model
import "time"
type Attendance struct{ID,EmployeeID int;CheckIn time.Time;CheckOut *time.Time;AttendanceDate time.Time}
type Leave struct{ID,EmployeeID int;LeaveDate time.Time;Reason,Status string}
