package service

import (
	"context"
	"configurable-employee-api/models"
	"configurable-employee-api/repository"
)

type EmployeeService interface {
	ListEmployees(ctx context.Context) ([]models.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) ListEmployees(ctx context.Context) ([]models.Employee, error) {
	return s.repo.GetAll(ctx)
}