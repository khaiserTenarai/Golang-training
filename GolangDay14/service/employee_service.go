package service

import (
	"employee_management/model"
	"employee_management/repository"
)

type EmployeeService struct {
	Repo repository.EmployeeRepository
}

func (s *EmployeeService) GetAll(department string) ([]model.Employee, error) {
	return s.Repo.FindAll(department)
}

func (s *EmployeeService) GetByID(id string) (model.Employee, error) {
	return s.Repo.FindByID(id)
}

func (s *EmployeeService) Create(employee *model.Employee) error {
	return s.Repo.Create(employee)
}

func (s *EmployeeService) Update(employee *model.Employee) error {
	return s.Repo.Update(employee)
}

func (s *EmployeeService) Delete(id string) error {
	return s.Repo.Delete(id)
}