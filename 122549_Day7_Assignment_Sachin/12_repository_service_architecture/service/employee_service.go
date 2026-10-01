package service

import (
	"context"
	"example.com/q12-repository-service/model"
)

type EmployeeService interface {
	Add(context.Context, model.Employee) error
	List(context.Context) ([]model.Employee, error)
}
