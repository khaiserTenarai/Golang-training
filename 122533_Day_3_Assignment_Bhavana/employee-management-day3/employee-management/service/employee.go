
package service

import (
	"errors"
	"fmt"
	"strings"

	"employee-management/model"
	"employee-management/utility"
)

var ErrEmployeeNotFound = errors.New("employee not found")
var ErrDuplicateEmployee = errors.New("employee already exists")


type EmployeeService struct {
	employees []model.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (s *EmployeeService) AddEmployee(emp model.Employee) error {
	emp.Name = utility.CleanString(emp.Name)
	emp.Email = strings.ToLower(utility.CleanString(emp.Email))

	if err := utility.ValidateName(emp.Name); err != nil {
		return fmt.Errorf("could not add employee: %w", err)
	}
	if err := utility.ValidateEmail(emp.Email); err != nil {
		return fmt.Errorf("could not add employee: %w", err)
	}
	if err := utility.ValidateAge(emp.Age); err != nil {
		return fmt.Errorf("could not add employee: %w", err)
	}
	if err := utility.ValidateSalary(emp.Salary); err != nil {
		return fmt.Errorf("could not add employee: %w", err)
	}

	_, err := s.GetEmployee(emp.ID)
	if err == nil {
		return fmt.Errorf("could not add employee: %w", ErrDuplicateEmployee)
	}
	if !errors.Is(err, ErrEmployeeNotFound) {
		return err
	}

	s.employees = append(s.employees, emp)
	return nil
}

func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {
	for _, e := range s.employees {
		if e.ID == id {
			return e, nil
		}
	}
	return model.Employee{}, fmt.Errorf("get employee failed: %w", ErrEmployeeNotFound)
}

func (s *EmployeeService) ListEmployees() []model.Employee {
	return s.employees
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	for i, e := range s.employees {
		if e.ID == id {
			s.employees = append(s.employees[:i], s.employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("delete employee failed: %w", ErrEmployeeNotFound)
}
