package service

import (
	"errors"

	"employee-app/model"
	"employee-app/repository"
)

type EmployeeService interface {
	CreateEmployee(id int, name, dept string, salary float64) (model.Employee, error)
	FetchEmployee(id int) (model.Employee, error)
	FetchAllEmployees() []model.Employee
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(id int, name, dept string, salary float64) (model.Employee, error) {
	if name == "" {
		return model.Employee{}, errors.New("invalid name")
	}
	if salary < 0 {
		return model.Employee{}, errors.New("salary cannot be negative")
	}

	emp := model.Employee{
		ID:         id,
		Name:       name,
		Department: dept,
		Salary:     salary,
	}

	err := s.repo.Save(emp)
	if err != nil {
		return model.Employee{}, err
	}
	return emp, nil
}

func (s *employeeService) FetchEmployee(id int) (model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) FetchAllEmployees() []model.Employee {
	return s.repo.GetAll()
}
