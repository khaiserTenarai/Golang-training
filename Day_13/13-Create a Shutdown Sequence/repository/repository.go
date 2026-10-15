package repository

import (
	"context"

	"employee-api/model"

	"gorm.io/gorm"
)

// EmployeeRepository defines the database operations used by the service.
type EmployeeRepository interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id uint) (*model.Employee, error)
}

// employeeRepository implements EmployeeRepository using GORM.
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
	// Pass the request context to the database operation.
	return r.db.WithContext(ctx).Create(employee).Error
}

// GetByID retrieves an employee by ID.
func (r *employeeRepository) GetByID(
	ctx context.Context,
	id uint,
) (*model.Employee, error) {
	var employee model.Employee

	if err := r.db.WithContext(ctx).First(&employee, id).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}