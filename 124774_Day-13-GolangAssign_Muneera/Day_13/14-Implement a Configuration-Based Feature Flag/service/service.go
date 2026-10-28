package service

import (
	"context"

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
func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

// Create validates and saves an employee.
func (s *employeeService) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	if err := validation.ValidateEmployee(employee); err != nil {
		return err
	}

	return s.repo.Create(ctx, employee)
}

// GetByID retrieves an employee by ID.
func (s *employeeService) GetByID(
	ctx context.Context,
	id uint,
) (*model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}