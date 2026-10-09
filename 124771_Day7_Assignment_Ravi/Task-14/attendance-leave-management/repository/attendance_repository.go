package repository

import (
	"time"

	"attendance-leave-management/model"
)

type AttendanceRepository interface {
	GetEmployees() ([]model.Employee, error)

	CheckIn(employeeID int, date time.Time, checkIn time.Time) error
	CheckOut(employeeID int, date time.Time, checkOut time.Time) error

	GetAttendanceReport(
		employeeID int,
		from time.Time,
		to time.Time,
	) ([]model.Attendance, error)

	ApplyLeave(leave model.Leave) error
	GetLeaves(employeeID int) ([]model.Leave, error)
	GetAllPendingLeaves() ([]model.Leave, error)
	UpdateLeaveStatus(leaveID int, status string) error
}
