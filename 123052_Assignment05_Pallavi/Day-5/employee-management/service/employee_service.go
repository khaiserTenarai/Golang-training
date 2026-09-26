package service

import (
	"errors"

	"employee-management/model"
	"employee-management/repository"
)

// Error when employee is not found
var ErrEmployeeNotFound = errors.New("employee not found")

// EmployeeService contains business logic
type EmployeeService struct {
	repository repository.EmployeeRepository
}

// Create service
func NewEmployeeService(
	repository repository.EmployeeRepository,
) *EmployeeService {

	return &EmployeeService{
		repository: repository,
	}
}

// Add employee
func (s *EmployeeService) AddEmployee(employee model.Employee) {

	s.repository.AddEmployee(employee)
}

// Get employee
func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {

	employee, found := s.repository.GetEmployee(id)

	if !found {
		return model.Employee{}, ErrEmployeeNotFound
	}

	return employee, nil
}

// Delete employee
func (s *EmployeeService) DeleteEmployee(id int) error {

	deleted := s.repository.DeleteEmployee(id)

	if !deleted {
		return ErrEmployeeNotFound
	}

	return nil
}

// Get all employees
func (s *EmployeeService) GetAllEmployees() []model.Employee {

	return s.repository.GetAllEmployees()
}