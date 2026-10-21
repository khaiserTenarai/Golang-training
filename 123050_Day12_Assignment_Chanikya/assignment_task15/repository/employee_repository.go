package repository

import (
	"context"
	"employee_api/model"
	"sync"
)

type EmployeeRepository interface {
	Insert(ctx context.Context, emp model.Employee) model.Employee
	FindAll(ctx context.Context) []model.Employee
}

type InMemoryEmployeeRepository struct {
	mu        sync.RWMutex
	employees []model.Employee
	currentID int
}

func NewEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make([]model.Employee, 0),
		currentID: 0,
	}
}

func (r *InMemoryEmployeeRepository) Insert(ctx context.Context, emp model.Employee) model.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Handle simulated multi-threaded execution safely
	r.currentID++
	emp.ID = r.currentID
	r.employees = append(r.employees, emp)
	return emp
}

func (r *InMemoryEmployeeRepository) FindAll(ctx context.Context) []model.Employee {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return structural copy to safeguard slice access
	list := make([]model.Employee, len(r.employees))
	copy(list, r.employees)
	return list
}
