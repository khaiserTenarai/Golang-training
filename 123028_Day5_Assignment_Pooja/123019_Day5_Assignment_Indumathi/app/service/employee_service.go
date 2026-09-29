package service

import (
	"errors"

	"app/model"
	"app/repository"
)

type EmployeeService interface {
	RegisterEmployee(emp *model.Employee) error
	GetEmployee(id int) (*model.Employee, error)
	GetAllEmployees() ([]model.Employee, error)
	GiveRaise(id int, percent float64) error
	DeactivateEmployee(id int) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) RegisterEmployee(emp *model.Employee) error {
	if emp.Salary <= 0 {
		return errors.New("business rule violation: salary must be greater than zero")
	}
	return s.repo.Send(emp)
}

func (s *employeeService) GetEmployee(id int) (*model.Employee, error) {
	return s.repo.Receive(id)
}

func (s *employeeService) GetAllEmployees() ([]model.Employee, error) {
	return s.repo.ReceiveAll()
}

func (s *employeeService) GiveRaise(id int, percent float64) error {
	if percent <= 0 {
		return errors.New("business rule violation: raise percent must be positive")
	}

	emp, err := s.repo.Receive(id)
	if err != nil {
		return err
	}

	emp.Salary += emp.Salary * (percent / 100.0)

	return s.repo.Send(emp)
}

func (s *employeeService) DeactivateEmployee(id int) error {
	emp, err := s.repo.Receive(id)
	if err != nil {
		return err
	}

	emp.IsActive = false

	return s.repo.Send(emp)
}