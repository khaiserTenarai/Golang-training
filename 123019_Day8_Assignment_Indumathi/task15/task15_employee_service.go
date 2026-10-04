package main

import (
	"errors"
	"sync"
)

type EmployeeService struct {
	mu        sync.RWMutex
	employees map[int]Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]Employee),
	}
}

func (s *EmployeeService) AddEmployee(e Employee) error {
	if err := ValidateEmployee(e); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.employees[e.ID] = e
	return nil
}

func (s *EmployeeService) GetEmployee(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	e, exists := s.employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	return e, nil
}