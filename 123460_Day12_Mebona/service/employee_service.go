package service

import (
	"context"
	"errors"

	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, emp *model.Employee) error
	GetEmployees(ctx context.Context, filter model.EmployeeFilter) ([]model.Employee, int, error)
	CheckHealth(ctx context.Context) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CheckHealth(ctx context.Context) error {
	return s.repo.Ping(ctx)
}

func (s *employeeService) CreateEmployee(ctx context.Context, emp *model.Employee) error {
	// Validation Layer
	if emp.Name == "" {
		return errors.New("validation error: name is required")
	}
	if emp.Email == "" {
		return errors.New("validation error: email is required")
	}
	if emp.Salary <= 0 {
		return errors.New("validation error: salary must be greater than zero")
	}
	return s.repo.Create(ctx, emp)
}

func (s *employeeService) GetEmployees(ctx context.Context, filter model.EmployeeFilter) ([]model.Employee, int, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 || filter.Limit > 100 {
		filter.Limit = 10
	}
	return s.repo.GetAll(ctx, filter)
}