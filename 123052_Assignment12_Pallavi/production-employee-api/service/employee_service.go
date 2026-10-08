package service

import (
	"errors"
	"production-employee-api/models"
	"production-employee-api/repository"
	"strings"
)

type EmployeeService interface {
	FetchEmployees(params models.QueryParams) ([]models.Employee, int)
	FetchByID(id string) (*models.Employee, error)
	CreateEmployee(emp models.Employee) (*models.Employee, error)
	CheckReadiness() bool
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CheckReadiness() bool {
	return s.repo.IsReady()
}

func (s *employeeService) FetchEmployees(params models.QueryParams) ([]models.Employee, int) {
	return s.repo.GetAll(params)
}

func (s *employeeService) FetchByID(id string) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) CreateEmployee(emp models.Employee) (*models.Employee, error) {
	// Validation Layer Logic
	if strings.TrimSpace(emp.Name) == "" {
		return nil, errors.New("validation failed: name is required")
	}
	if !strings.Contains(emp.Email, "@") {
		return nil, errors.New("validation failed: valid email is required")
	}
	created := s.repo.Create(emp)
	return &created, nil
}