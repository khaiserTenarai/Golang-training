// Package service contains business rules.
package service

import (
	"context"
	"employee-management/model"
	"employee-management/repository"
	"errors"
)

// EmployeeService defines employee business operations.
type EmployeeService interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string) ([]model.Employee, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
}

// employeeService implements EmployeeService.
type employeeService struct{ repository repository.EmployeeRepository }

// NewEmployeeService injects the repository.
func NewEmployeeService(r repository.EmployeeRepository) EmployeeService { return &employeeService{r} }

func (s *employeeService) Create(ctx context.Context, e *model.Employee) error {
	if err := validate(e); err != nil {
		return err
	}
	return s.repository.Create(ctx, e)
}
func (s *employeeService) GetAll(ctx context.Context, sort string) ([]model.Employee, error) {
	return s.repository.GetAll(ctx, sort)
}
func (s *employeeService) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than zero")
	}
	return s.repository.GetByID(ctx, id)
}
func (s *employeeService) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than zero")
	}
	if err := validate(e); err != nil {
		return nil, err
	}
	return s.repository.Update(ctx, id, e)
}
func (s *employeeService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("employee id must be greater than zero")
	}
	return s.repository.Delete(ctx, id)
}

func validate(e *model.Employee) error {
	if e == nil {
		return errors.New("employee body is required")
	}
	if e.Name == "" {
		return errors.New("employee name is required")
	}
	if e.Email == "" {
		return errors.New("employee email is required")
	}
	if e.Department == "" {
		return errors.New("employee department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
