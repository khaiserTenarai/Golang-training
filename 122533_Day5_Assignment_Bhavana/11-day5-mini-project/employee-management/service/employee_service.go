package service

import (
	"fmt"

	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	AddEmployee(emp model.Employee) error
	GetEmployee(id int) (model.Employee, error)
	DeleteEmployee(id int) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) AddEmployee(emp model.Employee) error {
	if emp.Name == "" {
		return fmt.Errorf("add employee failed: name cannot be empty")
	}
	s.repo.Add(emp)
	return nil
}

func (s *employeeService) GetEmployee(id int) (model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) DeleteEmployee(id int) error {
	return s.repo.Delete(id)
}
