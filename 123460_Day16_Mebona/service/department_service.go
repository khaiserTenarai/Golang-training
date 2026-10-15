package service

import (
	"fmt"

	"employee-management/apperror"
	"employee-management/model"
	"employee-management/repository"
)

type DepartmentService interface {
	List() ([]*model.Department, error)
	Get(id uint) (*model.Department, error)
	Employees(departmentID uint) ([]*model.Employee, error)
}

type departmentService struct {
	depts     repository.DepartmentRepository
	employees repository.EmployeeRepository
}

func NewDepartmentService(d repository.DepartmentRepository, e repository.EmployeeRepository) DepartmentService {
	return &departmentService{depts: d, employees: e}
}

func (s *departmentService) List() ([]*model.Department, error) { return s.depts.List() }

func (s *departmentService) Get(id uint) (*model.Department, error) {
	d, err := s.depts.GetByID(id)
	if isNotFound(err) {
		return nil, apperror.NotFound(fmt.Sprintf("department with id %d not found", id))
	}
	return d, err
}

func (s *departmentService) Employees(departmentID uint) ([]*model.Employee, error) {
	return s.employees.List(model.EmployeeFilter{DepartmentID: &departmentID})
}
