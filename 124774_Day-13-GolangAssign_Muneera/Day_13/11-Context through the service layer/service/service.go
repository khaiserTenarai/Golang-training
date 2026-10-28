package service

import (
	"context"
	"errors"
	"fmt"

	"employee-api/model"
	"employee-api/repository"
	"employee-api/validation"
)

// EmployeeService defines employee business operations.
type EmployeeService interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id uint) (*model.Employee, error)
}

// employeeService implements EmployeeService.
type employeeService struct {
	repo repository.EmployeeRepository
}

// NewEmployeeService creates an employee service.
func NewEmployeeService(
	repo repository.EmployeeRepository,
) EmployeeService {
	return &employeeService{repo: repo}
}

// Create validates and saves an employee.
func (s *employeeService) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	// Business input validation happens in the service layer.
	if err := validation.ValidateEmployee(employee); err != nil {
		return err
	}

	// Send valid data to the repository.
	if err := s.repo.Create(ctx, employee); err != nil {
		return fmt.Errorf("create employee: %w", err)
	}

	return nil
}

// GetByID retrieves an employee.
func (s *employeeService) GetByID(
	ctx context.Context,
	id uint,
) (*model.Employee, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	employee, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get employee: %w", err)
	}

	return employee, nil
}

// IsValidationError checks whether an error is a known validation error.
func IsValidationError(err error) bool {
	return errors.Is(err, validation.ErrNameRequired) ||
		errors.Is(err, validation.ErrEmailRequired) ||
		errors.Is(err, validation.ErrEmailInvalid)
}