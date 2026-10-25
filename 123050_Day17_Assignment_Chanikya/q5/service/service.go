package service

import (
	"errors"

	"employee-management/model"
	"employee-management/repository"
	"employee-management/utility"

	"gorm.io/gorm"
)

type EmployeeService interface {
	Create(employee *model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(id uint) (*model.Employee, error)
	Update(id uint, employee model.Employee) (*model.Employee, error)
	Delete(id uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

// NewEmployeeService creates the employee service.
func NewEmployeeService(
	repo repository.EmployeeRepository,
) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) Create(employee *model.Employee) error {
	if err := utility.ValidateEmployee(employee); err != nil {
		return err
	}

	return s.repo.Create(employee)
}

func (s *employeeService) GetAll() ([]model.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) GetByID(id uint) (*model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) Update(
	id uint,
	employee model.Employee,
) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("employee ID must be positive")
	}

	if err := utility.ValidateEmployee(&employee); err != nil {
		return nil, err
	}

	return s.repo.Update(id, employee)
}

func (s *employeeService) Delete(id uint) error {
	if id == 0 {
		return errors.New("employee ID must be positive")
	}

	return s.repo.Delete(id)
}

// IsNotFound checks GORM's not-found error, including wrapped errors.
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
