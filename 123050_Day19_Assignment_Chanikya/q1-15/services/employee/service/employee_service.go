package service

import (
	"context"
	"errors"

	"employee-management-microservices/services/employee/entity"
	"employee-management-microservices/services/employee/repository"
	"employee-management-microservices/shared/httpclient"

	"gorm.io/gorm"
)

type EmployeeService struct {
	Repo             *repository.EmployeeRepository
	DepartmentClient *httpclient.DepartmentClient
}

func NewEmployeeService(
	repo *repository.EmployeeRepository,
	departmentClient *httpclient.DepartmentClient,
) *EmployeeService {
	return &EmployeeService{
		Repo:             repo,
		DepartmentClient: departmentClient,
	}
}

func (s *EmployeeService) validateDepartment(
	ctx context.Context,
	departmentID uint,
	correlationID string,
) error {
	_, err := s.DepartmentClient.GetByID(
		ctx,
		departmentID,
		correlationID,
	)
	return err
}

func (s *EmployeeService) Create(
	ctx context.Context,
	correlationID string,
	name, email string,
	salary float64,
	departmentID uint,
) (*entity.Employee, error) {
	if err := s.validateDepartment(ctx, departmentID, correlationID); err != nil {
		return nil, err
	}

	employee := &entity.Employee{
		Name:         name,
		Email:        email,
		Salary:       salary,
		DepartmentID: departmentID,
	}

	if err := s.Repo.Create(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *EmployeeService) GetAll() ([]entity.Employee, error) {
	return s.Repo.GetAll()
}

func (s *EmployeeService) GetByID(id uint) (*entity.Employee, error) {
	return s.Repo.GetByID(id)
}

func (s *EmployeeService) Update(
	ctx context.Context,
	correlationID string,
	id uint,
	name, email string,
	salary float64,
	departmentID uint,
) (*entity.Employee, error) {
	employee, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if err := s.validateDepartment(ctx, departmentID, correlationID); err != nil {
		return nil, err
	}

	employee.Name = name
	employee.Email = email
	employee.Salary = salary
	employee.DepartmentID = departmentID

	if err := s.Repo.Update(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *EmployeeService) Delete(id uint) error {
	err := s.Repo.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return err
}
