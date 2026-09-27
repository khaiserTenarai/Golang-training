package service

import (
	"example.com/employee-management/model"
	"example.com/employee-management/repository"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repository: repository}
}

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) {
	s.repository.Save(employee)
}

func (s *EmployeeServiceImpl) GetAllEmployees() []model.Employee {
	return s.repository.FindAll()
}
