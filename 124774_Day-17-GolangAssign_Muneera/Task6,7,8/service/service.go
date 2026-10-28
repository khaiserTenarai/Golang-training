package service

import (
	"errors"

	"employee-management/model"
	"employee-management/repository"
	"employee-management/utility"

	"gorm.io/gorm"
)

type EmployeeService interface {
	Create(e *model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(id uint) (*model.Employee, error)
	Update(id uint, e *model.Employee) (*model.Employee, error)
	Delete(id uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: r}
}

func (s *employeeService) Create(e *model.Employee) error {
	if err := utility.ValidateEmployee(e); err != nil {
		return err
	}

	return s.repo.Create(e)
}

func (s *employeeService) GetAll() ([]model.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) GetByID(id uint) (*model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) Update(
	id uint,
	e *model.Employee,
) (*model.Employee, error) {
	if err := utility.ValidateEmployee(e); err != nil {
		return nil, err
	}

	return s.repo.Update(id, e)
}

func (s *employeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
