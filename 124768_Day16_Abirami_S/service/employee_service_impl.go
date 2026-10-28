package service

import (
	"errors"
	"fmt"
	"strings"

	"day16-employee-graphql/model"
	"day16-employee-graphql/repository"
)

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetEmployee(id int) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than 0")
	}
	return s.repo.GetByID(id)
}

func (s *employeeService) GetEmployees() ([]model.Employee, error) {
	return s.repo.GetAll()
}

func validateEmployee(employee model.Employee) error {
	if strings.TrimSpace(employee.Name) == "" {
		return errors.New("employee name is required")
	}
	if !strings.Contains(employee.Email, "@") {
		return errors.New("employee email must contain @")
	}
	if employee.Age < 18 || employee.Age > 100 {
		return errors.New("employee age must be between 18 and 100")
	}
	if employee.Salary < 0 {
		return errors.New("employee salary cannot be negative")
	}
	if employee.DepartmentID <= 0 {
		return errors.New("departmentId must be greater than 0")
	}
	return nil
}

func (s *employeeService) CreateEmployee(employee model.Employee) (*model.Employee, error) {
	if err := validateEmployee(employee); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	return s.repo.Create(employee)
}

func (s *employeeService) UpdateEmployee(id int, employee model.Employee) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than 0")
	}
	if err := validateEmployee(employee); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	return s.repo.Update(id, employee)
}

func (s *employeeService) DeleteEmployee(id int) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than 0")
	}
	return s.repo.Delete(id)
}
