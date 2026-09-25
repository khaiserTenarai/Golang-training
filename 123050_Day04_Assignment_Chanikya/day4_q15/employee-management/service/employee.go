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
	employees []model.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: []model.Employee{},
	}
}

func (s *EmployeeService) AddEmployee(emp model.Employee) error {

	emp.Name = utility.CleanString(emp.Name)
	emp.Email = utility.CleanString(emp.Email)

	if err := utility.ValidateName(emp.Name); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := utility.ValidateEmail(emp.Email); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := utility.ValidateAge(emp.Age); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := utility.ValidateSalary(emp.Salary); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	for _, employee := range s.employees {
		if employee.ID == emp.ID {
			return fmt.Errorf(
				"cannot add employee %d: %w",
				emp.ID,
				ErrDuplicateEmployee,
			)
		}
	}

	s.employees = append(s.employees, emp)

	return nil
}

func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {

	for _, employee := range s.employees {
		if employee.ID == id {
			return employee, nil
		}
	}

	return model.Employee{}, fmt.Errorf(
		"employee %d: %w",
		id,
		ErrEmployeeNotFound,
	)
}

func (s *EmployeeService) DeleteEmployee(id int) error {

	for i, employee := range s.employees {
		if employee.ID == id {

			s.employees = append(
				s.employees[:i],
				s.employees[i+1:]...,
			)

			return nil
		}
	}

	return fmt.Errorf(
		"employee %d: %w",
		id,
		ErrEmployeeNotFound,
	)
}
