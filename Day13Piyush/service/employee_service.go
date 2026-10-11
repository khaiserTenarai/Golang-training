// Task 11: Pass context through services
// Context flows: Handler → Service → Repository
// This enables cancellation and timeout propagation.

package service

import (
	"context"
	"errors"
	"strings"

	"configurable_api/models"
	"configurable_api/repository"
)

type EmployeeService interface {
	Create(ctx context.Context, emp models.Employee) (models.Employee, error)
	GetAll(ctx context.Context, params models.QueryParams) ([]models.Employee, int, error)
	GetByID(ctx context.Context, id int) (models.Employee, error)
	Update(ctx context.Context, id int, emp models.Employee) (models.Employee, error)
	Delete(ctx context.Context, id int) error
	Count() int
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) Create(ctx context.Context, emp models.Employee) (models.Employee, error) {
	if err := validate(emp); err != nil {
		return models.Employee{}, err
	}
	return s.repo.Create(ctx, emp) // Task 11: context passed through
}

func (s *employeeService) GetAll(ctx context.Context, params models.QueryParams) ([]models.Employee, int, error) {
	return s.repo.GetAll(ctx, params)
}

func (s *employeeService) GetByID(ctx context.Context, id int) (models.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *employeeService) Update(ctx context.Context, id int, emp models.Employee) (models.Employee, error) {
	if err := validate(emp); err != nil {
		return models.Employee{}, err
	}
	return s.repo.Update(ctx, id, emp)
}

func (s *employeeService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *employeeService) Count() int {
	return s.repo.Count()
}

func validate(emp models.Employee) error {
	if strings.TrimSpace(emp.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(emp.Email, "@") {
		return errors.New("valid email required")
	}
	if emp.Age < 18 || emp.Age > 65 {
		return errors.New("age must be 18-65")
	}
	if emp.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
