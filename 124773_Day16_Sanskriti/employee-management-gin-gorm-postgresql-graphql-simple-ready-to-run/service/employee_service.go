package service

import (
	"employee-management/model"
	"employee-management/repository"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type EmployeeService interface {
	Create(*model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}
type employeeService struct{ repo repository.EmployeeRepository }

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService { return &employeeService{r} }
func validate(e *model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)
	if e.Name == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(e.Email, "@") {
		return errors.New("valid email is required")
	}
	if e.Department == "" {
		return errors.New("department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
func (s *employeeService) Create(e *model.Employee) error {
	if err := validate(e); err != nil {
		return err
	}
	return s.repo.Create(e)
}
func (s *employeeService) GetAll() ([]model.Employee, error)        { return s.repo.GetAll() }
func (s *employeeService) GetByID(id uint) (*model.Employee, error) { return s.repo.GetByID(id) }
func (s *employeeService) Update(id uint, e *model.Employee) (*model.Employee, error) {
	if err := validate(e); err != nil {
		return nil, err
	}
	return s.repo.Update(id, e)
}
func (s *employeeService) Delete(id uint) error { return s.repo.Delete(id) }
func IsNotFound(err error) bool                 { return errors.Is(err, gorm.ErrRecordNotFound) }
