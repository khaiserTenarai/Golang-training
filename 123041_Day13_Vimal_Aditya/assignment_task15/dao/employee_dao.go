package dao

import (
	"context"

	"employeeapi/model"
)

type EmployeeDAO interface {
	Save(ctx context.Context, employee model.Employee) (model.Employee, error)
	FindByID(ctx context.Context, id int) (model.Employee, error)
	FindAll(ctx context.Context) ([]model.Employee, error)
	Delete(ctx context.Context, id int) error
}
