package service
import (
	"employee-management/model"
	"employee-management/repository"
	"employee-management/utility"
)

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

func (s *EmployeeServiceImpl) AddEmployee(
	employee model.Employee,
) error {

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSal(employee.Salary); err != nil {
		return err
	}

	return s.repository.AddEmployee(employee)
}

func (s *EmployeeServiceImpl) GetAllEmployees() []model.Employee {

	return s.repository.GetAllEmployees()
}

func (s *EmployeeServiceImpl) GetEmployeeByID(
	id int,
) model.Employee {

	return s.repository.GetEmployeeByID(id)
}

func (s *EmployeeServiceImpl) UpdateEmployee(
	employee model.Employee,
) error {

	// Validate updated employee data

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSal(employee.Salary); err != nil {
		return err
	}

	// Check whether employee exists

	existingEmployee := s.repository.GetEmployeeByID(employee.ID)

	if existingEmployee.ID == 0 {
		return utility.ErrEmployeeNotFound
	}

	return s.repository.UpdateEmployee(employee)
}

func (s *EmployeeServiceImpl) DeleteEmployee(
	id int,
) error {

	// Check whether employee exists

	employee := s.repository.GetEmployeeByID(id)

	if employee.ID == 0 {
		return utility.ErrEmployeeNotFound
	}

	return s.repository.DeleteEmployee(id)
}