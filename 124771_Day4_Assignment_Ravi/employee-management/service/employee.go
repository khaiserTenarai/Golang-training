package service

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)

// Sentinel errors.
var ErrEmployeeNotFound = errors.New(
	"employee not found",
)

var ErrDuplicateEmployee = errors.New(
	"employee already exists",
)

type EmployeeService struct {
	employees map[int]model.Employee
}

// NewEmployeeService creates a new employee service.
func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]model.Employee),
	}
}

// AddEmployee adds an employee.
func (s *EmployeeService) AddEmployee(
	employee model.Employee,
) error {

	// Clean string fields.
	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	// Validate employee.
	if err := utility.ValidateEmployee(employee); err != nil {
		return fmt.Errorf(
			"failed to add employee: %w",
			err,
		)
	}

	// Check duplicate ID.
	if _, exists := s.employees[employee.ID]; exists {
		return fmt.Errorf(
			"failed to add employee with ID %d: %w",
			employee.ID,
			ErrDuplicateEmployee,
		)
	}

	s.employees[employee.ID] = employee

	return nil
}

// GetEmployee retrieves an employee by ID.
func (s *EmployeeService) GetEmployee(
	id int,
) (model.Employee, error) {

	employee, exists := s.employees[id]

	if !exists {
		return model.Employee{}, fmt.Errorf(
			"failed to get employee with ID %d: %w",
			id,
			ErrEmployeeNotFound,
		)
	}

	return employee, nil
}

// DeleteEmployee deletes an employee by ID.
func (s *EmployeeService) DeleteEmployee(
	id int,
) error {

	if _, exists := s.employees[id]; !exists {
		return fmt.Errorf(
			"failed to delete employee with ID %d: %w",
			id,
			ErrEmployeeNotFound,
		)
	}

	delete(s.employees, id)

	return nil
}
