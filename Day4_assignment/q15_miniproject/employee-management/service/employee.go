package service

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)


var ErrEmployeeNotFound = errors.New("employee not found")
var ErrDuplicateEmployee = errors.New("employee already exists")


type EmployeeService struct {
	employees map[int]model.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{employees: make(map[int]model.Employee)}
}

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


func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {
	e, found := s.employees[id]
	if !found {
		return model.Employee{}, fmt.Errorf("GetEmployee(%d): %w", id, ErrEmployeeNotFound)
	}
	return e, nil
}


func (s *EmployeeService) DeleteEmployee(id int) error {
	if _, found := s.employees[id]; !found {
		return fmt.Errorf("DeleteEmployee(%d): %w", id, ErrEmployeeNotFound)
	}
	delete(s.employees, id)
	return nil
}
