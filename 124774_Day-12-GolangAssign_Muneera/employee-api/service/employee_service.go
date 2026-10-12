package service

import (
	"context"
	"day12/model"
)

type EmployeeService interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string, string, string, string, int, int) (*model.EmployeeListResponse, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
}
