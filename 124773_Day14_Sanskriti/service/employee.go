package service

import (
	"context"

	"employee-api/model"
	"employee-api/repository"
	"employee-api/validation"
)

type EmployeeService interface {
	List(context.Context, string, string, string, int, int) ([]model.Employee, error)
	Get(context.Context, int64) (model.Employee, error)
	Create(context.Context, model.Employee) (model.Employee, error)
	Update(context.Context, int64, model.Employee) (model.Employee, error)
	Delete(context.Context, int64) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) List(ctx context.Context, department, sortBy, order string, limit, offset int) ([]model.Employee, error) {
	return s.repo.List(ctx, department, sortBy, order, limit, offset)
}

func (s *employeeService) Get(ctx context.Context, id int64) (model.Employee, error) {
	return s.repo.Get(ctx, id)
}

func (s *employeeService) Create(ctx context.Context, e model.Employee) (model.Employee, error) {
	if err := validation.ValidateEmployee(e); err != nil {
		return model.Employee{}, err
	}
	return s.repo.Create(ctx, e)
}

func (s *employeeService) Update(ctx context.Context, id int64, e model.Employee) (model.Employee, error) {
	if err := validation.ValidateEmployee(e); err != nil {
		return model.Employee{}, err
	}
	return s.repo.Update(ctx, id, e)
}

func (s *employeeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
