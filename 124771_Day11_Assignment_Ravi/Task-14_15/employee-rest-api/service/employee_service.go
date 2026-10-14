package service

import (
	"context"
	"errors"
	"strings"

	"employee-rest-api/model"
)

type EmployeeStore interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context) ([]model.Employee, error)
	GetByID(context.Context, int64) (model.Employee, error)
	Update(context.Context, int64, *model.Employee) (model.Employee, error)
	Delete(context.Context, int64) error
}

type EmployeeService struct {
	repo EmployeeStore
}

func NewEmployeeService(repo EmployeeStore) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(ctx context.Context, e *model.Employee) error {
	if err := validate(e); err != nil {
		return err
	}
	normalize(e)
	return s.repo.Create(ctx, e)
}

func (s *EmployeeService) GetAll(ctx context.Context) ([]model.Employee, error) {
	return s.repo.GetAll(ctx)
}

func (s *EmployeeService) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) Update(ctx context.Context, id int64, e *model.Employee) (model.Employee, error) {
	if err := validate(e); err != nil {
		return model.Employee{}, err
	}
	normalize(e)
	return s.repo.Update(ctx, id, e)
}

func (s *EmployeeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func normalize(e *model.Employee) {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)
}

func validate(e *model.Employee) error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(e.Email) == "" || !strings.Contains(strings.TrimSpace(e.Email), "@") {
		return errors.New("a valid email is required")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
