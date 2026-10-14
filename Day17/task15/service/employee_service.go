package service

import (
	"fmt"

	"employee-platform/models"
	"employee-platform/repository"
)

type EmployeeService interface {
	ListEmployees(filter models.EmployeeFilter, page models.Pagination, sortBy string, sortOrder string) ([]models.Employee, int, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	CreateEmployee(emp *models.Employee) error
	GetDepartmentByID(id string) (*models.Department, error)
	BatchGetDepartments(ids []string) (map[string]models.Department, error)
}

type employeeService struct {
	repo repository.Repository
}

func NewEmployeeService(repo repository.Repository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) ListEmployees(filter models.EmployeeFilter, page models.Pagination, sortBy string, sortOrder string) ([]models.Employee, int, error) {
	return s.repo.GetEmployees(filter, page, sortBy, sortOrder)
}

func (s *employeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *employeeService) CreateEmployee(emp *models.Employee) error {
	// Input validation
	if emp.Name == "" || emp.Salary <= 0 {
		return fmt.Errorf("invalid employee details")
	}
	return s.repo.CreateEmployee(emp)
}

func (s *employeeService) GetDepartmentByID(id string) (*models.Department, error) {
	return s.repo.GetDepartmentByID(id)
}

func (s *employeeService) BatchGetDepartments(ids []string) (map[string]models.Department, error) {
	return s.repo.GetDepartmentsByIDs(ids)
}
