// 15. Day 8 Mini Project: Deliver a tested Employee Service with
// unit tests, coverage, formatting, static analysis, and race detection.
//
// This file is the service itself. service_test.go has the tests.
// main.go is a small interactive program that uses this service with
// real input from the user. README.md documents how to actually run all
// 5 quality checks against this folder.

package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// EmployeeService stores employees in memory. The mutex is what makes it
// safe to call Add/Get/Delete from multiple goroutines at the same time -
// see the TestConcurrentAdd test and README.md for how that's checked.
type EmployeeService struct {
	mu        sync.Mutex
	employees map[int]Employee
	nextID    int
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]Employee),
		nextID:    1,
	}
}

func (s *EmployeeService) Add(name string, salary float64) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.employees[id] = Employee{ID: id, Name: name, Salary: salary}
	s.nextID++
	return id
}

func (s *EmployeeService) Get(id int) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	emp, ok := s.employees[id]
	if !ok {
		return Employee{}, fmt.Errorf("employee %d not found", id)
	}
	return emp, nil
}

func (s *EmployeeService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.employees[id]; !ok {
		return fmt.Errorf("employee %d not found", id)
	}
	delete(s.employees, id)
	return nil
}

func (s *EmployeeService) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.employees)
}
