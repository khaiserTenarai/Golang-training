package service

import (
	"api/models"
	"api/repository"
	"errors"
)

type EmployeeService interface {
	GetEmployees(page, limit int, deptFilter, sortBy string) []models.Employee
	CreateEmployee(emp models.Employee) (models.Employee, error)
}

type employeeServiceImpl struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo: repo}
}

func (s *employeeServiceImpl) GetEmployees(p, l int, dept, sort string) []models.Employee {
	if p < 1 { p = 1 }
	if l < 1 { l = 10 }
	return s.repo.GetEmployees(p, l, dept, sort)
}

func (s *employeeServiceImpl) CreateEmployee(emp models.Employee) (models.Employee, error) {
	// 9. Validation Layer[cite: 2]
	if emp.Name == "" || emp.Department == "" {
		return models.Employee{}, errors.New("name and department are required")
	}
	return s.repo.Save(emp), nil
}