package repository

import (
	"context"
	"day12/model"
)

type EmployeeRepository interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string, string, string, string, int, int) ([]model.Employee, int, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
}
