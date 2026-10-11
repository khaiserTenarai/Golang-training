// Package service implements business logic.
package service

import (
	"context"
	"errors"

	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	Create(ctx context.Context, emp *model.Employee) error
	GetAll(ctx context.Context, sort string) ([]model.Employee, error)
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	Update(ctx context.Context, id int64, emp *model.Employee) (*model.Employee, error)
	PartialUpdate(ctx context.Context, id int64, patch *model.PartialEmployee) (*model.Employee, error)
	Delete(ctx context.Context, id int64) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) Create(ctx context.Context, emp *model.Employee) error {
	if emp.Salary <= 0 {
		return errors.New("salary must be greater than zero")
	}
	return s.repo.Create(ctx, emp)
}

func (s *employeeService) GetAll(ctx context.Context, sort string) ([]model.Employee, error) {
	return s.repo.GetAll(ctx, sort)
}

func (s *employeeService) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *employeeService) Update(ctx context.Context, id int64, emp *model.Employee) (*model.Employee, error) {
	if emp.Salary <= 0 {
		return nil, errors.New("salary must be greater than zero")
	}
	return s.repo.Update(ctx, id, emp)
}

func (s *employeeService) PartialUpdate(ctx context.Context, id int64, patch *model.PartialEmployee) (*model.Employee, error) {
	if patch.Salary != nil && *patch.Salary <= 0 {
		return nil, errors.New("salary must be greater than zero")
	}
	return s.repo.PartialUpdate(ctx, id, patch)
}

func (s *employeeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}