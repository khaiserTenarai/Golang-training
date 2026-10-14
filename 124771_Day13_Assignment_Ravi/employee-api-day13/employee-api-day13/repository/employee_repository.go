package repository

import (
	"context"

	"employee-api-day13/model"
)

// EmployeeRepository is the persistence contract.
type EmployeeRepository interface {
	Create(ctx context.Context, input model.EmployeeInput) (model.Employee, error)
	List(ctx context.Context) ([]model.Employee, error)
	GetByID(ctx context.Context, id int64) (model.Employee, error)
	Update(ctx context.Context, id int64, input model.EmployeeInput) (model.Employee, error)
	Delete(ctx context.Context, id int64) error
}
