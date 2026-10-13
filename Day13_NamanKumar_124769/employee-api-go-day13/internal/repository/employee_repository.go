package repository

import (
	"context"
	"errors"
	"sort"
	"strings"
	"sync"
	"time"

	"employee-api-go/internal/domain"
)

var ErrNotFound = errors.New("employee not found")
var ErrEmailExists = errors.New("email already exists")

type EmployeeRepository interface {
	Create(ctx context.Context, input domain.CreateEmployeeInput) (domain.Employee, error)
	GetByID(ctx context.Context, id int64) (domain.Employee, error)
	Update(ctx context.Context, id int64, input domain.UpdateEmployeeInput) (domain.Employee, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) []domain.Employee
	Ready(ctx context.Context) bool
}

type MemoryEmployeeRepository struct {
	mu        sync.RWMutex
	employees map[int64]domain.Employee
	nextID    int64
}

func NewMemoryEmployeeRepository() *MemoryEmployeeRepository {
	return &MemoryEmployeeRepository{
		employees: make(map[int64]domain.Employee),
		nextID:    1,
	}
}

func (r *MemoryEmployeeRepository) Create(ctx context.Context, input domain.CreateEmployeeInput) (domain.Employee, error) {
	if err := ctx.Err(); err != nil {
		return domain.Employee{}, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.emailExists(input.Email, 0) {
		return domain.Employee{}, ErrEmailExists
	}

	employee := domain.Employee{
		ID:         r.nextID,
		Name:       strings.TrimSpace(input.Name),
		Email:      strings.ToLower(strings.TrimSpace(input.Email)),
		Department: strings.TrimSpace(input.Department),
		CreatedAt:  time.Now().UTC(),
	}
	r.employees[employee.ID] = employee
	r.nextID++
	return employee, nil
}

func (r *MemoryEmployeeRepository) GetByID(ctx context.Context, id int64) (domain.Employee, error) {
	if err := ctx.Err(); err != nil {
		return domain.Employee{}, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	employee, ok := r.employees[id]
	if !ok {
		return domain.Employee{}, ErrNotFound
	}
	return employee, nil
}

func (r *MemoryEmployeeRepository) Update(ctx context.Context, id int64, input domain.UpdateEmployeeInput) (domain.Employee, error) {
	if err := ctx.Err(); err != nil {
		return domain.Employee{}, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.employees[id]; !ok {
		return domain.Employee{}, ErrNotFound
	}
	if r.emailExists(input.Email, id) {
		return domain.Employee{}, ErrEmailExists
	}

	employee := r.employees[id]
	employee.Name = strings.TrimSpace(input.Name)
	employee.Email = strings.ToLower(strings.TrimSpace(input.Email))
	employee.Department = strings.TrimSpace(input.Department)
	r.employees[id] = employee
	return employee, nil
}

func (r *MemoryEmployeeRepository) Delete(ctx context.Context, id int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.employees[id]; !ok {
		return ErrNotFound
	}
	delete(r.employees, id)
	return nil
}

func (r *MemoryEmployeeRepository) List(ctx context.Context) []domain.Employee {
	if ctx.Err() != nil {
		return []domain.Employee{}
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Employee, 0, len(r.employees))
	for _, employee := range r.employees {
		result = append(result, employee)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func (r *MemoryEmployeeRepository) Ready(ctx context.Context) bool {
	if ctx.Err() != nil {
		return false
	}
	return r.employees != nil
}

func (r *MemoryEmployeeRepository) emailExists(email string, exceptID int64) bool {
	email = strings.ToLower(strings.TrimSpace(email))
	for id, employee := range r.employees {
		if id != exceptID && strings.EqualFold(employee.Email, email) {
			return true
		}
	}
	return false
}
