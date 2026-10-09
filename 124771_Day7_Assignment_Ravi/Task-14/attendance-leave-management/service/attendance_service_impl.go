package service

import (
	"errors"
	"strings"
	"time"

	"attendance-leave-management/model"
	"attendance-leave-management/repository"
)

type AttendanceServiceImpl struct {
	repository repository.AttendanceRepository
}

func NewAttendanceService(
	repository repository.AttendanceRepository,
) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		repository: repository,
	}
}

func (s *AttendanceServiceImpl) GetEmployees() ([]model.Employee, error) {
	return s.repository.GetEmployees()
}

func (s *AttendanceServiceImpl) CheckIn(employeeID int) error {
	if employeeID <= 0 {
		return errors.New("employee ID must be greater than zero")
	}

	now := time.Now()
	date := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	return s.repository.CheckIn(employeeID, date, now)
}

func (s *AttendanceServiceImpl) CheckOut(employeeID int) error {
	if employeeID <= 0 {
		return errors.New("employee ID must be greater than zero")
	}

	now := time.Now()
	date := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	return s.repository.CheckOut(employeeID, date, now)
}

func (s *AttendanceServiceImpl) GetAttendanceReport(
	employeeID int,
	from time.Time,
	to time.Time,
) ([]model.Attendance, error) {
	if employeeID <= 0 {
		return nil, errors.New("employee ID must be greater than zero")
	}

	if from.After(to) {
		return nil, errors.New("from date cannot be after to date")
	}

	return s.repository.GetAttendanceReport(
		employeeID,
		from,
		to,
	)
}

func (s *AttendanceServiceImpl) ApplyLeave(
	leave model.Leave,
) error {
	if leave.EmployeeID <= 0 {
		return errors.New("employee ID must be greater than zero")
	}

	if strings.TrimSpace(leave.Reason) == "" {
		return errors.New("leave reason cannot be empty")
	}

	if leave.LeaveDate.IsZero() {
		return errors.New("leave date is required")
	}

	return s.repository.ApplyLeave(leave)
}

func (s *AttendanceServiceImpl) GetLeaves(
	employeeID int,
) ([]model.Leave, error) {
	if employeeID <= 0 {
		return nil, errors.New("employee ID must be greater than zero")
	}

	return s.repository.GetLeaves(employeeID)
}

func (s *AttendanceServiceImpl) GetAllPendingLeaves() ([]model.Leave, error) {
	return s.repository.GetAllPendingLeaves()
}

func (s *AttendanceServiceImpl) ApproveLeave(
	leaveID int,
) error {
	if leaveID <= 0 {
		return errors.New("leave ID must be greater than zero")
	}

	return s.repository.UpdateLeaveStatus(
		leaveID,
		"APPROVED",
	)
}

func (s *AttendanceServiceImpl) RejectLeave(
	leaveID int,
) error {
	if leaveID <= 0 {
		return errors.New("leave ID must be greater than zero")
	}

	return s.repository.UpdateLeaveStatus(
		leaveID,
		"REJECTED",
	)
}
