package main

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	Create(e Employee) (Employee, error)
	GetByID(id int) (Employee, error)
	GetAll() []Employee
	Update(e Employee) error
	Delete(id int) error
}

type InMemoryEmployeeRepository struct {
	mu     sync.Mutex
	data   map[int]Employee
	nextID int
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		data:   make(map[int]Employee),
		nextID: 1,
	}
}

func (r *InMemoryEmployeeRepository) Create(e Employee) (Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	e.ID = r.nextID
	r.nextID++
	r.data[e.ID] = e
	return e, nil
}

func (r *InMemoryEmployeeRepository) GetByID(id int) (Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	e, ok := r.data[id]
	if !ok {
		return Employee{}, ErrNotFound
	}
	return e, nil
}

func (r *InMemoryEmployeeRepository) GetAll() []Employee {
	r.mu.Lock()
	defer r.mu.Unlock()
	result := make([]Employee, 0, len(r.data))
	for _, e := range r.data {
		result = append(result, e)
	}
	return result
}

func (r *InMemoryEmployeeRepository) Update(e Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[e.ID]; !ok {
		return ErrNotFound
	}
	r.data[e.ID] = e
	return nil
}

func (r *InMemoryEmployeeRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return ErrNotFound
	}
	delete(r.data, id)
	return nil
}
