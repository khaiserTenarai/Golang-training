package service

import (
	"context"
	"errors"
	"strings"

	"employee-management-microservices/services/employee/client"
	"employee-management-microservices/services/employee/model"
	"employee-management-microservices/services/employee/repository"
	"employee-management-microservices/utility"

	"gorm.io/gorm"
)

var (
	ErrInvalidInput       = errors.New("invalid employee input")
	ErrDepartmentNotFound = errors.New("department not found")
)

// Service contains employee business rules.
type Service struct {
	repo       repository.Repository
	department *client.DepartmentClient
}

func New(
	repo repository.Repository,
	department *client.DepartmentClient,
) *Service {
	return &Service{
		repo:       repo,
		department: department,
	}
}

func (s *Service) Create(
	ctx context.Context,
	e *model.Employee,
) error {
	if err := utility.ValidateEmployee(
		e.Name, e.Email, e.DepartmentID, e.Salary,
	); err != nil {
		return errors.Join(ErrInvalidInput, err)
	}

	// Check that the department exists before saving.
	exists, err := s.department.Exists(ctx, e.DepartmentID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrDepartmentNotFound
	}

	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)

	return s.repo.Create(e)
}

func (s *Service) GetAll() ([]model.Employee, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(
	ctx context.Context,
	e *model.Employee,
) error {
	if err := utility.ValidateEmployee(
		e.Name, e.Email, e.DepartmentID, e.Salary,
	); err != nil {
		return errors.Join(ErrInvalidInput, err)
	}

	// Check the department before changing the employee.
	exists, err := s.department.Exists(ctx, e.DepartmentID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrDepartmentNotFound
	}

	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)

	return s.repo.Update(e)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}

// IsNotFound checks for a missing employee.
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
