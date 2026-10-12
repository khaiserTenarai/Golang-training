package repository

import (
	"sync"

	"layered/model"
)

// The repository is the only layer that knows where the data lives.
type EmployeeRepository struct {
	mu        sync.Mutex
	employees []model.Employee
	nextID    int
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{nextID: 1}
}

func (r *EmployeeRepository) Save(e model.Employee) model.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()

	e.ID = r.nextID
	r.nextID++
	r.employees = append(r.employees, e)
	return e
}

func (r *EmployeeRepository) FindAll() []model.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]model.Employee, len(r.employees))
	copy(result, r.employees)
	return result
}

func (r *EmployeeRepository) FindByID(id int) (model.Employee, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, e := range r.employees {
		if e.ID == id {
			return e, true
		}
	}
	return model.Employee{}, false
}
