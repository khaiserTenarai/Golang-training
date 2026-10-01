package model

type Employee struct {
	ID   int64
	Name string
}
type Attendance struct {
	ID           int64
	EmployeeID   int64
	EmployeeName string
	CheckIn      string
	CheckOut     string
}
type Leave struct {
	ID           int64
	EmployeeID   int64
	EmployeeName string
	LeaveDate    string
	Reason       string
	Status       string
}
