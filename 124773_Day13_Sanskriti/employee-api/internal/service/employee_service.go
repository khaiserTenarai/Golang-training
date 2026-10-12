package service

import (
	"context"
	"log"

	"employee-api/internal/model"
	"employee-api/internal/repository"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, emp *model.Employee) error
	GetEmployee(ctx context.Context, id string) (*model.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		repo: repo,
	}
}

func (s *employeeService) CreateEmployee(ctx context.Context, emp *model.Employee) error {
	log.Printf("[SERVICE] Creating new employee record for Email: %s", emp.Email)
	return s.repo.Create(ctx, emp)
}

func (s *employeeService) GetEmployee(ctx context.Context, id string) (*model.Employee, error) {
	log.Printf("[SERVICE] Fetching employee profile for ID: %s", id)
	return s.repo.GetByID(ctx, id)
}
