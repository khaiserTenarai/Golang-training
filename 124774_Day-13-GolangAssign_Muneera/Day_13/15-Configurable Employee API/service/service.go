package service

import (
	"context"

	"employee-api/model"
	"employee-api/repository"
	"employee-api/validation"
)

type EmployeeService interface {
	Create(
		ctx context.Context,
		employee *model.Employee,
	) error

	GetByID(
		ctx context.Context,
		id uint,
	) (*model.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

// NewEmployeeService creates the service.
func NewEmployeeService(
	repo repository.EmployeeRepository,
) EmployeeService {
	return &employeeService{repo: repo}
}

// Create validates the employee before saving.
func (s *employeeService) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	if err := validation.ValidateEmployee(employee); err != nil {
		return err
	}

	return s.repo.Create(ctx, employee)
}

// GetByID retrieves an employee.
func (s *employeeService) GetByID(
	ctx context.Context,
	id uint,
) (*model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}