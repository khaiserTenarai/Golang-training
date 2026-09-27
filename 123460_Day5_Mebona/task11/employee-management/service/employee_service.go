package service

import (
	"employee-management/model"
	"employee-management/repository"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	repository repository.EmployeeRepository,
) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		repository: repository,
	}
}

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) bool {
	_, exists := s.repository.Get(employee.ID)

	if exists {
		return false
	}

	s.repository.Add(employee)
	return true
}

func (s *EmployeeServiceImpl) GetEmployee(
	id int,
) (model.Employee, bool) {

	return s.repository.Get(id)
}

func (s *EmployeeServiceImpl) GetAllEmployees() []model.Employee {
	return s.repository.GetAll()
}

func (s *EmployeeServiceImpl) UpdateEmployee(
	employee model.Employee,
) bool {

	return s.repository.Update(employee)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) bool {
	return s.repository.Delete(id)
}