package service

import (
	"errors"

	"attendance-leave/model"
	"attendance-leave/repository"
	"attendance-leave/utility"
)

type AttendanceServiceImpl struct {
	repository repository.AttendanceRepository
}

func NewAttendanceService(
	repository repository.AttendanceRepository,
) AttendanceService {

	return &AttendanceServiceImpl{
		repository: repository,
	}
}

func (s *AttendanceServiceImpl) AddEmployee(
	employee model.Employee,
) error {

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	return s.repository.AddEmployee(employee)
}

func (s *AttendanceServiceImpl) CheckIn(
	employeeID int,
) error {

	if err := utility.ValidateID(employeeID); err != nil {
		return err
	}

	return s.repository.CheckIn(employeeID)
}

func (s *AttendanceServiceImpl) CheckOut(
	employeeID int,
) error {

	if err := utility.ValidateID(employeeID); err != nil {
		return err
	}

	return s.repository.CheckOut(employeeID)
}

func (s *AttendanceServiceImpl) GetAttendance() []model.Attendance {

	return s.repository.GetAttendance()
}

func (s *AttendanceServiceImpl) ApplyLeave(
	leave model.Leave,
) error {

	if err := utility.ValidateID(leave.EmployeeID); err != nil {
		return err
	}

	if err := utility.ValidateReason(leave.Reason); err != nil {
		return err
	}

	return s.repository.ApplyLeave(leave)
}

func (s *AttendanceServiceImpl) ApproveLeave(
	leaveID int,
) error {

	if leaveID <= 0 {
		return errors.New("leave ID must be greater than 0")
	}

	return s.repository.UpdateLeaveStatus(
		leaveID,
		"Approved",
	)
}

func (s *AttendanceServiceImpl) RejectLeave(
	leaveID int,
) error {

	if leaveID <= 0 {
		return errors.New("leave ID must be greater than 0")
	}

	return s.repository.UpdateLeaveStatus(
		leaveID,
		"Rejected",
	)
}

func (s *AttendanceServiceImpl) GetLeaves() []model.Leave {

	return s.repository.GetLeaves()
}
