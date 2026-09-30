package service
import (
	"attendance_leave/model"
	"attendance_leave/repository"
	"attendance_leave/utility"
)

type EmployeeService interface {
	Save(employee model.Employee) error
	FindByID(id int) (model.Employee, error)
	FindAll() ([]model.Employee, error)
}

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	repository repository.EmployeeRepository,
) EmployeeService {

	return &EmployeeServiceImpl{
		repository: repository,
	}
}

func (s *EmployeeServiceImpl) Save(
	employee model.Employee,
) error {

	err := utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(employee.Email)

	if err != nil {
		return err
	}

	return s.repository.Save(employee)
}

func (s *EmployeeServiceImpl) FindByID(
	id int,
) (model.Employee, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return model.Employee{}, err
	}

	return s.repository.FindByID(id)
}

func (s *EmployeeServiceImpl) FindAll() (
	[]model.Employee,
	error,
) {

	return s.repository.FindAll()
}
