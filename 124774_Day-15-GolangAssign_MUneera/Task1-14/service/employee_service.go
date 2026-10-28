package service

import (
	"context"
	"errors"
	"strings"

	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, int, int, string, string) ([]model.Employee, int, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
}

type employeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	r repository.EmployeeRepository,
) EmployeeService {
	return &employeeService{repository: r}
}

func validate(e *model.Employee) error {
	if e == nil {
		return errors.New("employee body is required")
	}

	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)

	if e.Name == "" {
		return errors.New("employee name is required")
	}

	if !strings.Contains(e.Email, "@") {
		return errors.New("valid employee email is required")
	}

	if e.Department == "" {
		return errors.New("employee department is required")
	}

	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return nil
}

func (s *employeeService) Create(
	ctx context.Context,
	e *model.Employee,
) error {
	if err := validate(e); err != nil {
		return err
	}

	return s.repository.Create(ctx, e)
}

func (s *employeeService) GetAll(
	ctx context.Context,
	page, size int,
	department, sortBy string,
) ([]model.Employee, int, error) {
	return s.repository.GetAll(
		ctx,
		page,
		size,
		department,
		sortBy,
	)
}

func (s *employeeService) GetByID(
	ctx context.Context,
	id int64,
) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than zero")
	}

	return s.repository.GetByID(ctx, id)
}

func (s *employeeService) Update(
	ctx context.Context,
	id int64,
	e *model.Employee,
) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than zero")
	}

	if err := validate(e); err != nil {
		return nil, err
	}

	return s.repository.Update(ctx, id, e)
}

func (s *employeeService) Delete(
	ctx context.Context,
	id int64,
) error {
	if id <= 0 {
		return errors.New("employee id must be greater than zero")
	}

	return s.repository.Delete(ctx, id)
}
