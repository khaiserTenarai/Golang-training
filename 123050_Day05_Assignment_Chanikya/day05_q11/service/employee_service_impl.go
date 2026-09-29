package service

import (
	"employee-management/model"
	"employee-management/repository"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{
		repository: repo,
	}
}

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) {
	s.repository.Add(employee)
}

func (s *EmployeeServiceImpl) GetEmployee(id int) (model.Employee, bool) {
	return s.repository.GetByID(id)
}

func (s *EmployeeServiceImpl) GetAllEmployees() []model.Employee {
	return s.repository.GetAll()
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) {
	s.repository.Delete(id)
}
