package service
import (
	"attendance_leave/model"
	"attendance_leave/repository"
	"attendance_leave/utility"
)

type LeaveService interface {
	Apply(leave model.Leave) error
	Approve(id int) error
	Reject(id int) error
	FindByEmployee(employeeID int) ([]model.LeaveReport, error)
	FindAll() ([]model.LeaveReport, error)
}

type LeaveServiceImpl struct {
	repository repository.LeaveRepository
}

func NewLeaveService(
	repository repository.LeaveRepository,
) LeaveService {

	return &LeaveServiceImpl{
		repository: repository,
	}
}

func (s *LeaveServiceImpl) Apply(
	leave model.Leave,
) error {

	err := utility.ValidateLeave(leave)

	if err != nil {
		return err
	}

	return s.repository.Apply(leave)
}

func (s *LeaveServiceImpl) Approve(
	id int,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Approve(id)
}

func (s *LeaveServiceImpl) Reject(
	id int,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Reject(id)
}

func (s *LeaveServiceImpl) FindByEmployee(
	employeeID int,
) ([]model.LeaveReport, error) {

	err := utility.ValidateID(employeeID)

	if err != nil {
		return nil, err
	}

	return s.repository.FindByEmployee(employeeID)
}

func (s *LeaveServiceImpl) FindAll() (
	[]model.LeaveReport,
	error,
) {

	return s.repository.FindAll()
}
