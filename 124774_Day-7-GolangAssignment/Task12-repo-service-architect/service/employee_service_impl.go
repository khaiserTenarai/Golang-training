package service

import (
	"errors"

	"employee-app/model"
	"employee-app/repository"
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

	if employee.Name == "" {
		return errors.New("name cannot be empty")
	}

	if employee.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return s.repository.Save(employee)
}

func (s *EmployeeServiceImpl) GetEmployees() []model.Employee {

	return s.repository.FindAll()
}
