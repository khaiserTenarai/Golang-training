package service

import (
	"errors"
	"strings"

	"layered/model"
	"layered/repository"
)

var (
	ErrNotFound = errors.New("employee not found")
	ErrInvalid  = errors.New("a name and a salary above 0 are required")
)

// The service holds the business rules - it knows nothing about HTTP.
type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(e model.Employee) (model.Employee, error) {
	if strings.TrimSpace(e.Name) == "" || e.Salary <= 0 {
		return model.Employee{}, ErrInvalid
	}
	return s.repo.Save(e), nil
}

func (s *EmployeeService) List() []model.Employee {
	return s.repo.FindAll()
}

func (s *EmployeeService) Get(id int) (model.Employee, error) {
	e, found := s.repo.FindByID(id)
	if !found {
		return model.Employee{}, ErrNotFound
	}
	return e, nil
}
