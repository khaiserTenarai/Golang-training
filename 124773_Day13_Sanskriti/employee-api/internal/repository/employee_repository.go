package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"employee-api/internal/database"
	"employee-api/internal/model"
)

type EmployeeRepository interface {
	Create(ctx context.Context, emp *model.Employee) error
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}

type memoryRepo struct {
	mu   sync.RWMutex
	db   *database.DB
	data map[string]*model.Employee
}

func NewEmployeeRepository(db *database.DB) EmployeeRepository {
	return &memoryRepo{
		db:   db,
		data: make(map[string]*model.Employee),
	}
}

func (r *memoryRepo) Create(ctx context.Context, emp *model.Employee) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.mu.Lock()
		defer r.mu.Unlock()

		emp.CreatedAt = time.Now()
		r.data[emp.ID] = emp
		return nil
	}
}

func (r *memoryRepo) GetByID(ctx context.Context, id string) (*model.Employee, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		r.mu.RLock()
		defer r.mu.RUnlock()

		emp, found := r.data[id]
		if !found {
			return nil, fmt.Errorf("employee with id %s not found", id)
		}
		return emp, nil
	}
}
