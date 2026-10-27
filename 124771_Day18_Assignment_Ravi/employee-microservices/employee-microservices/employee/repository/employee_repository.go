package repository

import (
	"errors"
	"sync"

	"employee-microservices/employee/model"
)

// EmployeeRepository is a small in-memory repository for learning.
// Task 4: repository layer. Task 7: replace with a service-owned database later.
type EmployeeRepository struct {
	mu     sync.Mutex
	nextID int
	items  map[int]model.Employee
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{nextID: 1, items: make(map[int]model.Employee)}
}

func (r *EmployeeRepository) List() []model.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()
	result := make([]model.Employee, 0, len(r.items))
	for _, e := range r.items {
		result = append(result, e)
	}
	return result
}

func (r *EmployeeRepository) Create(e model.Employee) model.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()
	e.ID = r.nextID
	r.nextID++
	r.items[e.ID] = e
	return e
}

func (r *EmployeeRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return errors.New("employee not found")
	}
	delete(r.items, id)
	return nil
}
