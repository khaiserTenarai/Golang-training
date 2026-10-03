package service

import (
	"assignment_task12/model"

	"assignment_task12/repository"
)

type EmployeeServiceImpl struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(
	repo repository.EmployeeRepository,
) EmployeeService {

	return &EmployeeServiceImpl{

		repo: repo,
	}

}

func (s *EmployeeServiceImpl) CreateEmployee(
	employee model.Employee,
) error {

	return s.repo.Create(employee)

}

func (s *EmployeeServiceImpl) GetEmployees() ([]model.Employee, error) {

	return s.repo.GetAll()

}
