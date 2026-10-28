package repository

import (
	"employee-department-microservices/department-service/model"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("department not found")

type DepartmentRepository struct {
	mu     sync.RWMutex
	nextID int
	items  map[int]model.Department
}

func NewDepartmentRepository() *DepartmentRepository {
	r := &DepartmentRepository{nextID: 1, items: make(map[int]model.Department)}
	r.items[1] = model.Department{ID: 1, Name: "IT", Description: "Information Technology"}
	r.nextID = 2
	r.items[2] = model.Department{ID: 2, Name: "HR", Description: "Human Resources"}
	r.nextID = 3
	return r
}
func (r *DepartmentRepository) List() []model.Department {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := []model.Department{}
	for _, v := range r.items {
		out = append(out, v)
	}
	return out
}
func (r *DepartmentRepository) Get(id int) (model.Department, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	if !ok {
		return model.Department{}, ErrNotFound
	}
	return v, nil
}
func (r *DepartmentRepository) Create(v model.Department) model.Department {
	r.mu.Lock()
	defer r.mu.Unlock()
	v.ID = r.nextID
	r.nextID++
	r.items[v.ID] = v
	return v
}
func (r *DepartmentRepository) Update(id int, v model.Department) (model.Department, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return model.Department{}, ErrNotFound
	}
	v.ID = id
	r.items[id] = v
	return v, nil
}
func (r *DepartmentRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	return nil
}
