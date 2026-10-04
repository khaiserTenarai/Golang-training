package service

import "attendance-leave/model"

type AttendanceService interface {
	AddEmployee(employee model.Employee) error

	CheckIn(employeeID int) error

	CheckOut(employeeID int) error

	GetAttendance() []model.Attendance

	ApplyLeave(leave model.Leave) error

	ApproveLeave(leaveID int) error

	RejectLeave(leaveID int) error

	GetLeaves() []model.Leave
}
