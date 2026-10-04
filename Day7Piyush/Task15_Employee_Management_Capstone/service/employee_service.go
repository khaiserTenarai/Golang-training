package service

import (
	"errors"
	"task15_employee_management_capstone/models"
	"task15_employee_management_capstone/repository"
)

type EmployeeService interface {
	CreateEmployee(emp models.Employee) (models.Employee, error)
	GetAllEmployees(page, pageSize int, search, sortBy, sortOrder string) ([]models.Employee, int, error)
	GetEmployeeByID(id int) (models.Employee, error)
	UpdateEmployee(emp models.Employee) (models.Employee, error)
	DeleteEmployee(id int) error
	UpdateSalary(empID int, newSalary float64, reason string) error
	GetSalaryHistory(empID int) ([]models.SalaryHistory, error)
}

type EmployeeServiceImpl struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repo: repo}
}

func (s *EmployeeServiceImpl) CreateEmployee(emp models.Employee) (models.Employee, error) {
	if emp.Name == "" {
		return models.Employee{}, errors.New("name is required")
	}
	if emp.Email == "" {
		return models.Employee{}, errors.New("email is required")
	}
	if emp.Salary < 0 {
		return models.Employee{}, errors.New("salary cannot be negative")
	}
	return s.repo.Create(emp)
}

func (s *EmployeeServiceImpl) GetAllEmployees(page, pageSize int, search, sortBy, sortOrder string) ([]models.Employee, int, error) {
	return s.repo.GetAll(page, pageSize, search, sortBy, sortOrder)
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
		return models.Employee{}, errors.New("name is required")
	}
	return s.repo.Update(emp)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) error {
	if id <= 0 {
		return errors.New("invalid employee ID")
	}
	return s.repo.Delete(id)
}

func (s *EmployeeServiceImpl) UpdateSalary(empID int, newSalary float64, reason string) error {
	if empID <= 0 {
		return errors.New("invalid employee ID")
	}
	if newSalary < 0 {
		return errors.New("salary cannot be negative")
	}
	return s.repo.UpdateSalary(empID, newSalary, reason)
}

func (s *EmployeeServiceImpl) GetSalaryHistory(empID int) ([]models.SalaryHistory, error) {
	return s.repo.GetSalaryHistory(empID)
}
