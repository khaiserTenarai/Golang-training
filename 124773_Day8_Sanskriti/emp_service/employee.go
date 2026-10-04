package empservice
package main

import "errors"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type EmployeeService struct {
	employees []Employee
}

// AddEmployee adds a new employee.
func (s *EmployeeService) AddEmployee(employee Employee) error {
	if employee.ID <= 0 {
		return errors.New("invalid employee ID")
	}

	if employee.Name == "" {
		return errors.New("employee name cannot be empty")
	}

	if employee.Age <= 0 {
		return errors.New("employee age must be greater than 0")
	}

	if employee.Salary <= 0 {
		return errors.New("employee salary must be greater than 0")
	}

	s.employees = append(s.employees, employee)

	return nil
}

// GetEmployee returns an employee by ID.
func (s *EmployeeService) GetEmployee(id int) (Employee, error) {
	for _, employee := range s.employees {
		if employee.ID == id {
			return employee, nil
		}
	}

	return Employee{}, errors.New("employee not found")
}

// GetAllEmployees returns all employees.
func (s *EmployeeService) GetAllEmployees() []Employee {
	return s.employees
}