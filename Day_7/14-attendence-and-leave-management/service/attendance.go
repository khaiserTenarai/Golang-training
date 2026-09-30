package service
import (
	"attendance_leave/model"
	"attendance_leave/repository"
	"attendance_leave/utility"
)

type AttendanceService interface {
	CheckIn(employeeID int) error
	CheckOut(employeeID int) error
	FindByEmployee(employeeID int) ([]model.AttendanceReport, error)
	FindAll() ([]model.AttendanceReport, error)
}

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

func (s *AttendanceServiceImpl) CheckIn(
	employeeID int,
) error {

	err := utility.ValidateID(employeeID)

	if err != nil {
		return err
	}

	return s.repository.CheckIn(employeeID)
}

func (s *AttendanceServiceImpl) CheckOut(
	employeeID int,
) error {

	err := utility.ValidateID(employeeID)

	if err != nil {
		return err
	}

	return s.repository.CheckOut(employeeID)
}

func (s *AttendanceServiceImpl) FindByEmployee(
	employeeID int,
) ([]model.AttendanceReport, error) {

	err := utility.ValidateID(employeeID)

	if err != nil {
		return nil, err
	}

	return s.repository.FindByEmployee(employeeID)
}

func (s *AttendanceServiceImpl) FindAll() (
	[]model.AttendanceReport,
	error,
) {

	return s.repository.FindAll()
}
