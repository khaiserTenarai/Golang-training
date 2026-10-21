package repository

import (
	"context"
	"day13_project/model"
	"fmt"
	"sync"
)

type EmployeeRepository struct {
	mu        sync.RWMutex
	employees map[string]model.Employee
}

func NewEmployeeRepository() *EmployeeRepository {
	repo := &EmployeeRepository{
		employees: make(map[string]model.Employee),
	}
	// Inject seed records mock data matching production scenarios
	repo.employees["E01"] = model.Employee{ID: "E01", Name: "Alice Smith", Role: "admin", Salary: 95000}
	repo.employees["E02"] = model.Employee{ID: "E02", Name: "Bob Jones", Role: "engineer", Salary: 82000}
	return repo
}

func (r *EmployeeRepository) GetAll(ctx context.Context) ([]model.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var list []model.Employee
	for _, emp := range r.employees {
		list = append(list, emp)
	}
	return list, nil
}

func (r *EmployeeRepository) Save(ctx context.Context, emp model.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.employees[emp.ID] = emp
	return nil
}
