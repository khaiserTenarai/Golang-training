package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

var errNotFound = errors.New("employee not found")

// the mutex makes the store safe when several requests arrive together
type Store struct {
	mu        sync.Mutex
	employees []Employee
	nextID    int
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (s *Store) Add(emp Employee) Employee {
	s.mu.Lock()
	defer s.mu.Unlock()

	emp.ID = s.nextID
	s.nextID++
	s.employees = append(s.employees, emp)
	return emp
}

// List returns one page of employees plus the total number stored
func (s *Store) List(page, limit int) ([]Employee, int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	total := len(s.employees)

	start := (page - 1) * limit
	if start > total {
		start = total
	}
	end := start + limit
	if end > total {
		end = total
	}

	result := make([]Employee, end-start)
	copy(result, s.employees[start:end])
	return result, total
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, emp := range s.employees {
		if emp.ID == id {
			s.employees = append(s.employees[:i], s.employees[i+1:]...)
			return nil
		}
	}
	return errNotFound
}
