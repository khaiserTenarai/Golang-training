package service

import (
	"time"

	"attendance-leave-management/model"
)

type AttendanceService interface {
	GetEmployees() ([]model.Employee, error)

	CheckIn(employeeID int) error
	CheckOut(employeeID int) error

	GetAttendanceReport(
		employeeID int,
		from time.Time,
		to time.Time,
	) ([]model.Attendance, error)

	ApplyLeave(leave model.Leave) error
	GetLeaves(employeeID int) ([]model.Leave, error)
	GetAllPendingLeaves() ([]model.Leave, error)

	ApproveLeave(leaveID int) error
	RejectLeave(leaveID int) error
}
