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
	employees map[int]*model.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{employees: make(map[int]*model.Employee)}
}

func NewIDGenerator(start int) func() int {
	currentID := start
	return func() int {
		id := currentID
		currentID++
		return id
	}
}

func (s *EmployeeService) AddEmployee(emp *model.Employee) error {
	if _, exists := s.employees[emp.ID]; exists {
		return fmt.Errorf("add employee failed: %w", ErrDuplicateEmployee)
	}
	if err := utility.ValidateName(emp.Name); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}
	if err := utility.ValidateEmail(emp.Email); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}
	if err := utility.ValidateAge(emp.Age); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}
	if err := utility.ValidateSalary(emp.Salary); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}

	emp.Name = utility.CleanString(emp.Name)
	emp.Email = utility.CleanString(emp.Email)
	s.employees[emp.ID] = emp
	return nil
}

func (s *EmployeeService) GetEmployee(id int) (*model.Employee, error) {
	emp, exists := s.employees[id]
	if !exists {
		return nil, fmt.Errorf("get employee failed: %w", ErrEmployeeNotFound)
	}
	return emp, nil
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	if _, exists := s.employees[id]; !exists {
		return fmt.Errorf("delete employee failed: %w", ErrEmployeeNotFound)
	}
	delete(s.employees, id)
	return nil
}

func (s *EmployeeService) EmployeeStats() (int, float64) {
	count := 0
	total := 0.0
	for _, emp := range s.employees {
		count++
		total += emp.Salary
	}
	return count, total
}

func (s *EmployeeService) AllEmployees() []*model.Employee {
	result := make([]*model.Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		result = append(result, emp)
	}
	return result
}

func CalculateTotalSalary(salaries ...float64) float64 {
	total := 0.0
	for _, sal := range salaries {
		total += sal
	}
	return total
}
