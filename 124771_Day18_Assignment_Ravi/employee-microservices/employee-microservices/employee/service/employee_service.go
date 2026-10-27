package service

import (
	"errors"
	"strings"

	"employee-microservices/employee/model"
	"employee-microservices/employee/repository"
)

// EmployeeService contains business rules.
// Task 4: service layer. Task 12/13: Saga compensation can call DeleteCreatedEmployee.
type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) List() []model.Employee { return s.repo.List() }

func (s *EmployeeService) Create(e model.Employee) (model.Employee, error) {
	e.Name = strings.TrimSpace(e.Name)
	if e.Name == "" {
		return model.Employee{}, errors.New("name is required")
	}
	if e.DepartmentID <= 0 {
		return model.Employee{}, errors.New("departmentId must be greater than zero")
	}
	return s.repo.Create(e), nil
}

func (s *EmployeeService) DeleteCreatedEmployee(id int) error { return s.repo.Delete(id) }
