package service

import (
	"errors"

	"employee-management/model"
	"employee-management/utility"
)

type EmployeeService struct {
	employees []model.Employee
	nextID    int
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make([]model.Employee, 0),
		nextID:    1,
	}
}

func (s *EmployeeService) AddEmployee(name string, salary float64) (*model.Employee, error) {
	if err := utility.ValidateEmployee(name, salary); err != nil {
		return nil, err
	}

	emp := model.Employee{
		ID:     s.nextID,
		Name:   utility.FormatName(name),
		Salary: salary,
	}
	s.nextID++
	s.employees = append(s.employees, emp)
	return &emp, nil
}

func (s *EmployeeService) GetEmployee(id int) (*model.Employee, error) {
	for i := range s.employees {
		if s.employees[i].ID == id {
			return &s.employees[i], nil
		}
	}
	return nil, errors.New("employee not found")
}

func (s *EmployeeService) ListEmployees() []model.Employee {
	return s.employees
}
