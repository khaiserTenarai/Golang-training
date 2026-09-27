// Package service contains the business logic for managing employees -
// this is what main.go talks to, it doesn't touch model/utility directly.
package service

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)

// Sentinel errors - plain, comparable error values that callers can
// check for with errors.Is, even after they've been wrapped.
var ErrEmployeeNotFound = errors.New("employee not found")
var ErrDuplicateEmployee = errors.New("employee already exists")

// EmployeeService keeps employees in memory, keyed by ID.
type EmployeeService struct {
	employees map[int]model.Employee
}

// NewEmployeeService returns a ready-to-use, empty service.
func NewEmployeeService() *EmployeeService {
	return &EmployeeService{employees: make(map[int]model.Employee)}
}

// AddEmployee validates the given employee and stores it, unless an
// employee with that ID already exists.
func (s *EmployeeService) AddEmployee(e model.Employee) error {
	if _, exists := s.employees[e.ID]; exists {
		return fmt.Errorf("AddEmployee(%d): %w", e.ID, ErrDuplicateEmployee)
	}

	e.Name = utility.CleanString(e.Name)

	if err := utility.ValidateName(e.Name); err != nil {
		return fmt.Errorf("AddEmployee: %w", err)
	}
	if err := utility.ValidateEmail(e.Email); err != nil {
		return fmt.Errorf("AddEmployee: %w", err)
	}
	if err := utility.ValidateAge(e.Age); err != nil {
		return fmt.Errorf("AddEmployee: %w", err)
	}
	if err := utility.ValidateSalary(e.Salary); err != nil {
		return fmt.Errorf("AddEmployee: %w", err)
	}

	s.employees[e.ID] = e
	return nil
}

// GetEmployee looks up an employee by ID.
func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {
	e, found := s.employees[id]
	if !found {
		return model.Employee{}, fmt.Errorf("GetEmployee(%d): %w", id, ErrEmployeeNotFound)
	}
	return e, nil
}

// DeleteEmployee removes an employee by ID.
func (s *EmployeeService) DeleteEmployee(id int) error {
	if _, found := s.employees[id]; !found {
		return fmt.Errorf("DeleteEmployee(%d): %w", id, ErrEmployeeNotFound)
	}
	delete(s.employees, id)
	return nil
}
