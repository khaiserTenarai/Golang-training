package repository

import (
	"configurable-employee-api/models"
	"context"
	"sync"
)

type EmployeeRepository interface {
	GetAll(ctx context.Context) ([]models.Employee, error)
	Close() error
}

type memoryRepo struct {
	mu        sync.RWMutex
	employees map[string]models.Employee
}

func NewEmployeeRepository() EmployeeRepository {
	repo := &memoryRepo{
		employees: make(map[string]models.Employee),
	}
	repo.employees["1"] = models.Employee{ID: "1", Name: "Alice", Department: "Engineering", Salary: 90000}
	repo.employees["2"] = models.Employee{ID: "2", Name: "Bob", Department: "HR", Salary: 65000}
	return repo
}

func (r *memoryRepo) GetAll(ctx context.Context) ([]models.Employee, error) {
	r.RLock()
	defer r.RUnlock()

	var list []models.Employee
	for _, e := range r.employees {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			list = append(list, e)
		}
	}
	return list, nil
}

func (r *memoryRepo) RUnlock() {
	panic("unimplemented")
}

func (r *memoryRepo) RLock() {
	panic("unimplemented")
}

func (r *memoryRepo) Close() error {
	// Cleanup operations if needed
	return nil
}
