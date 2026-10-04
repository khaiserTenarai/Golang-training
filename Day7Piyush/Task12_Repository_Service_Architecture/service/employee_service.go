package service

import (
	"errors"
	"task12_repository_service_architecture/models"
	"task12_repository_service_architecture/repository"
)

// EmployeeService interface - defines the contract for business logic
type EmployeeService interface {
	CreateEmployee(emp models.Employee) (models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id int) (models.Employee, error)
	UpdateEmployee(emp models.Employee) (models.Employee, error)
	DeleteEmployee(id int) error
}

// EmployeeServiceImpl implements EmployeeService with business logic
type EmployeeServiceImpl struct {
	repo repository.EmployeeRepository // Depends on interface, not concrete type
}

// NewEmployeeService injects the repository dependency via interface
func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repo: repo}
}

func (s *EmployeeServiceImpl) CreateEmployee(emp models.Employee) (models.Employee, error) {
	if emp.Name == "" {
		return models.Employee{}, errors.New("employee name is required")
	}
	if emp.Email == "" {
		return models.Employee{}, errors.New("employee email is required")
	}
	if emp.Salary < 0 {
		return models.Employee{}, errors.New("salary cannot be negative")
	}
	return s.repo.Create(emp)
}

func (s *EmployeeServiceImpl) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.GetAll()
}

func (s *EmployeeServiceImpl) GetEmployeeByID(id int) (models.Employee, error) {
	if id <= 0 {
		return models.Employee{}, errors.New("invalid employee ID")
	}
	return s.repo.GetByID(id)
}

func (s *EmployeeServiceImpl) UpdateEmployee(emp models.Employee) (models.Employee, error) {
	if emp.ID <= 0 {
		return models.Employee{}, errors.New("invalid employee ID")
	}
	if emp.Name == "" {
		return models.Employee{}, errors.New("employee name is required")
	}
	return s.repo.Update(emp)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) error {
	if id <= 0 {
		return errors.New("invalid employee ID")
	}
	return s.repo.Delete(id)
}
