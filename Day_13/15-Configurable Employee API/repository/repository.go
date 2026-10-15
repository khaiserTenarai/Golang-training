package repository

import (
	"context"

	"employee-api/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(
		ctx context.Context,
		employee *model.Employee,
	) error

	GetByID(
		ctx context.Context,
		id uint,
	) (*model.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

// NewEmployeeRepository creates a repository.
func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

// Create saves an employee.
func (r *employeeRepository) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

// GetByID retrieves an employee by ID.
func (r *employeeRepository) GetByID(
	ctx context.Context,
	id uint,
) (*model.Employee, error) {
	var employee model.Employee

	err := r.db.WithContext(ctx).First(&employee, id).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}