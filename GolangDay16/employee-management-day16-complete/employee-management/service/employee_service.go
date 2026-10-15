package service

import (
	"employee-management/model"
	"employee-management/repository"
	"errors"
	"gorm.io/gorm"
	"net/mail"
	"strings"
)

type EmployeeService interface {
	Create(*model.Employee) error
	GetAll(model.EmployeeFilter) ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}
type employeeService struct {
	repo  repository.EmployeeRepository
	depts repository.DepartmentRepository
}

func NewEmployeeService(r repository.EmployeeRepository, d repository.DepartmentRepository) EmployeeService {
	return &employeeService{r, d}
}
func (s *employeeService) validate(e *model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	if e.Name == "" {
		return invalid("name is required")
	}
	if a, err := mail.ParseAddress(e.Email); err != nil || a.Address != e.Email {
		return invalid("valid email is required")
	}
	if e.Salary < 0 {
		return invalid("salary cannot be negative")
	}
	if e.DepartmentID == 0 {
		return invalid("departmentId is required")
	}
	if _, err := s.depts.GetByID(e.DepartmentID); err != nil {
		if IsNotFound(err) {
			return invalid("department does not exist")
		}
		return err
	}
	return nil
}
func saveError(err error) error {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrDuplicateEmail
	}
	return err
}
func (s *employeeService) Create(e *model.Employee) error {
	if err := s.validate(e); err != nil {
		return err
	}
	return saveError(s.repo.Create(e))
}
func (s *employeeService) GetAll(f model.EmployeeFilter) ([]model.Employee, error) {
	return s.repo.GetAll(f)
}
func (s *employeeService) GetByID(id uint) (*model.Employee, error) { return s.repo.GetByID(id) }
func (s *employeeService) Update(id uint, e *model.Employee) (*model.Employee, error) {
	if err := s.validate(e); err != nil {
		return nil, err
	}
	x, err := s.repo.Update(id, e)
	if err != nil {
		return nil, saveError(err)
	}
	return x, nil
}
func (s *employeeService) Delete(id uint) error { return s.repo.Delete(id) }
