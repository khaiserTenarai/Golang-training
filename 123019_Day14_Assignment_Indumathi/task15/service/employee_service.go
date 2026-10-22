package service

import (
	"employee-api/models"
	"employee-api/repository"
)

type EmployeeService interface {
	GetAllEmployees(department string) ([]models.Employee, error)
	GetEmployeeByID(id int) (*models.Employee, error)
	CreateEmployee(emp *models.Employee) error
	UpdateEmployee(id int, emp *models.Employee) error
	PatchEmployee(id int, update *models.UpdateEmployee) (*models.Employee, error)
	DeleteEmployee(id int) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAllEmployees(department string) ([]models.Employee, error) {
	return s.repo.GetAll(department)
}

func (s *employeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) CreateEmployee(emp *models.Employee) error {
	return s.repo.Create(emp)
}

func (s *employeeService) UpdateEmployee(id int, emp *models.Employee) error {
	return s.repo.Update(id, emp)
}

func (s *employeeService) PatchEmployee(id int, update *models.UpdateEmployee) (*models.Employee, error) {
	return s.repo.Patch(id, update)
}

func (s *employeeService) DeleteEmployee(id int) error {
	return s.repo.Delete(id)
}
