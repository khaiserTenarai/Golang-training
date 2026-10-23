package repository

import (
	"errors"
	"sync"

	"employee-department-microservices/employee-service/model"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository struct {
	mu     sync.RWMutex
	nextID int
	items  map[int]model.Employee
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{nextID: 1, items: make(map[int]model.Employee)}
}
func (r *EmployeeRepository) List() []model.Employee {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]model.Employee, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	return out
}
func (r *EmployeeRepository) Get(id int) (model.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	if !ok {
		return model.Employee{}, ErrNotFound
	}
	return v, nil
}
func (r *EmployeeRepository) Create(v model.Employee) model.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()
	v.ID = r.nextID
	r.nextID++
	r.items[v.ID] = v
	return v
}
func (r *EmployeeRepository) Update(id int, v model.Employee) (model.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return model.Employee{}, ErrNotFound
	}
	v.ID = id
	r.items[id] = v
	return v, nil
}
func (r *EmployeeRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	return nil
}
