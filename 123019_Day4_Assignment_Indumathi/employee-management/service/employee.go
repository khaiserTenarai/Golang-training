package service

import (
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)

type EmployeeService struct {
	employees []model.Employee
}

func (s *EmployeeService) AddEmployee(employee model.Employee) error {

	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	err := utility.ValidateName(employee.Name)
	if err != nil {
		return fmt.Errorf("cannot add employee: %w", err)
	}

	err = utility.ValidateEmail(employee.Email)
	if err != nil {
		return fmt.Errorf("cannot add employee: %w", err)
	}

	err = utility.ValidateAge(employee.Age)
	if err != nil {
		return fmt.Errorf("cannot add employee: %w", err)
	}

	err = utility.ValidateSalary(employee.Salary)
	if err != nil {
		return fmt.Errorf("cannot add employee: %w", err)
	}

	for _, existingEmployee := range s.employees {
		if existingEmployee.ID == employee.ID {
			return fmt.Errorf(
				"cannot add employee: %w",
				utility.ErrDuplicateEmployee,
			)
		}
	}

	s.employees = append(s.employees, employee)

	return nil
}

func (s *EmployeeService) GetEmployee(id int) (model.Employee, error) {

	for _, employee := range s.employees {
		if employee.ID == id {
			return employee, nil
		}
	}

	return model.Employee{}, fmt.Errorf(
		"cannot get employee: %w",
		utility.ErrEmployeeNotFound,
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
		"cannot delete employee: %w",
		utility.ErrEmployeeNotFound,
	)
}
