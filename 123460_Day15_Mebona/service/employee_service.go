package service

import (
	"context"
	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, emp *model.Employee) error
	GetEmployees(ctx context.Context, page, limit int) ([]model.Employee, int, error)
	GetEmployeeByID(ctx context.Context, id int64) (*model.Employee, error)
	UpdateAvatar(ctx context.Context, id int64, path string) error
	DeleteEmployee(ctx context.Context, id int64) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(ctx context.Context, emp *model.Employee) error {
	return s.repo.Save(ctx, emp)
}

func (s *employeeService) GetEmployees(ctx context.Context, page, limit int) ([]model.Employee, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.FindAll(ctx, offset, limit)
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, id int64) (*model.Employee, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *employeeService) UpdateAvatar(ctx context.Context, id int64, path string) error {
	return s.repo.UpdateAvatar(ctx, id, path)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}