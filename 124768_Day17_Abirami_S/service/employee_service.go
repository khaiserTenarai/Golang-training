package service

import (
	"errors"
	"math"

	"day17-employee-platform/model"
	"day17-employee-platform/repository"
)

type EmployeeService interface {
	GetAll(filter model.EmployeeFilter, sort model.SortOption, page, pageSize int) ([]model.Employee, int64, int, error)
	GetByID(id uint) (*model.Employee, error)
	Create(employee *model.Employee) error
	Update(id uint, employee *model.Employee) (*model.Employee, error)
	Delete(id uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAll(
	filter model.EmployeeFilter,
	sort model.SortOption,
	page, pageSize int,
) ([]model.Employee, int64, int, error) {

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	if pageSize > 100 {
		pageSize = 100
	}

	items, total, err := s.repo.GetAll(filter, sort, page, pageSize)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return items, total, totalPages, nil
}

func (s *employeeService) GetByID(id uint) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("employee id must be greater than zero")
	}
	return s.repo.GetByID(id)
}

func (s *employeeService) Create(employee *model.Employee) error {
	if err := validateEmployee(employee); err != nil {
		return err
	}
	return s.repo.Create(employee)
}

func (s *employeeService) Update(id uint, employee *model.Employee) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("employee id must be greater than zero")
	}
	if err := validateEmployee(employee); err != nil {
		return nil, err
	}
	return s.repo.Update(id, employee)
}

func (s *employeeService) Delete(id uint) error {
	if id == 0 {
		return errors.New("employee id must be greater than zero")
	}
	return s.repo.Delete(id)
}

func validateEmployee(employee *model.Employee) error {
	if employee == nil {
		return errors.New("employee input is required")
	}
	if employee.Name == "" {
		return errors.New("employee name is required")
	}
	if employee.Email == "" {
		return errors.New("employee email is required")
	}
	if employee.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	if employee.DepartmentID == 0 {
		return errors.New("department id is required")
	}
	return nil
}
