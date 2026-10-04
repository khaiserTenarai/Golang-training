package employee

import (
	"errors"
	"sync"
)

type Employee struct {
	ID         int
	Name       string
	Email      string
	Department string
	Salary     float64
}

type EmployeeService struct {
	mu        sync.RWMutex
	employees map[int]Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]Employee),
	}
}

func (s *EmployeeService) AddEmployee(employee Employee) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if employee.ID <= 0 {
		return errors.New("invalid employee ID")
	}

	if employee.Name == "" {
		return errors.New("employee name is required")
	}

	if employee.Email == "" {
		return errors.New("employee email is required")
	}

	if _, exists := s.employees[employee.ID]; exists {
		return errors.New("employee already exists")
	}

	s.employees[employee.ID] = employee
	return nil
}

func (s *EmployeeService) GetEmployee(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	employee, exists := s.employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}

	return employee, nil
}

func (s *EmployeeService) GetAllEmployees() []Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()

	employees := make([]Employee, 0, len(s.employees))

	for _, employee := range s.employees {
		employees = append(employees, employee)
	}

	return employees
}

func (s *EmployeeService) UpdateEmployee(employee Employee) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.employees[employee.ID]; !exists {
		return errors.New("employee not found")
	}

	s.employees[employee.ID] = employee
	return nil
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.employees[id]; !exists {
		return errors.New("employee not found")
	}

	delete(s.employees, id)
	return nil
}
