package service

import (
	"context"
	"errors"
	"strings"

	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	Create(ctx context.Context, e *model.Employee) error
	GetAll(ctx context.Context, sortBy string) ([]model.Employee, error)
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error)
	Delete(ctx context.Context, id int64) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func validate(e *model.Employee) error {
	if e == nil {
		return errors.New("employee payload is required")
	}
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department.Name = strings.TrimSpace(e.Department.Name)

	if e.Name == "" {
		return errors.New("employee name is required")
	}
	if !strings.Contains(e.Email, "@") {
		return errors.New("valid email address is required")
	}
	if e.Department.ID <= 0 || e.Department.Name == "" {
		return errors.New("valid department ID and name are required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}

func (s *employeeService) Create(ctx context.Context, e *model.Employee) error {
	if err := validate(e); err != nil {
		return err
	}
	return s.repo.Create(ctx, e)
}

func (s *employeeService) GetAll(ctx context.Context, sortBy string) ([]model.Employee, error) {
	return s.repo.GetAll(ctx, sortBy)
}

func (s *employeeService) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee ID must be greater than zero")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *employeeService) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee ID must be greater than zero")
	}
	if err := validate(e); err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, id, e)
}

func (s *employeeService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("employee ID must be greater than zero")
	}
	return s.repo.Delete(ctx, id)
}