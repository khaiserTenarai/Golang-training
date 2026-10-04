package repository

import "attendance-leave/model"

type AttendanceRepository interface {
	AddEmployee(employee model.Employee) error

	CheckIn(employeeID int) error

	CheckOut(employeeID int) error

	GetAttendance() []model.Attendance

	ApplyLeave(leave model.Leave) error

	UpdateLeaveStatus(
		leaveID int,
		status string,
	) error

	GetLeaves() []model.Leave
}
