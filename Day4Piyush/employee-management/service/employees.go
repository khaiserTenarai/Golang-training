package service

import (
	"employee-management/model"
	"employee-management/utility"
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("employee not found")
var ErrDuplicateEmployee = errors.New("employee already exists")

type EmployeeService struct {
	employees map[int]model.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]model.Employee),
	}
}

func (s *EmployeeService) AddEmployee(employee model.Employee) error {
	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}
	if err := utility.ValidateEmail(employee.Email); err != nil {
		return err
	}
	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}
	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}
	if _, exists := s.employees[employee.ID]; exists {
		return ErrDuplicateEmployee
	}
	s.employees[employee.ID] = employee

	return nil
}

func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {
	employee, exists := s.employees[id]

	if !exists {
		return model.Employee{}, fmt.Errorf(
			"failed to get employee %d: %w",
			id,
			ErrEmployeeNotFound,
		)
	}

	return employee, nil
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	if _, exists := s.employees[id]; !exists {
		return fmt.Errorf(
			"failed to delete employee %d: %w",
			id,
			ErrEmployeeNotFound,
		)
	}

	delete(s.employees, id)

	return nil
}

