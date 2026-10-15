package repository

import (
	"context"
	"errors"

	"employee-api/model"

	"gorm.io/gorm"
)

// EmployeeRepository defines employee database operations.
type EmployeeRepository interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id uint) (*model.Employee, error)
}

// employeeRepository implements EmployeeRepository.
type employeeRepository struct {
	db *gorm.DB
}

// NewEmployeeRepository creates a repository.
func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

// Create saves a new employee.
func (r *employeeRepository) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

// GetByID finds an employee by ID.
func (r *employeeRepository) GetByID(
	ctx context.Context,
	id uint,
) (*model.Employee, error) {
	var employee model.Employee

	err := r.db.WithContext(ctx).First(&employee, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &employee, nil
}