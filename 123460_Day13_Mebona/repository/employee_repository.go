package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"employee-api/model"
)

var ErrClosed = errors.New("repository is closed")

// EmployeeRepository is the storage contract. Every method takes a context, so a cancelled
// request or a timeout stops the work. A PostgreSQL version would implement the same interface.
type EmployeeRepository interface {
	Create(ctx context.Context, e model.Employee) (model.Employee, error)
	GetByID(ctx context.Context, id int) (model.Employee, error)
	List(ctx context.Context, f model.ListFilter) (items []model.Employee, total int, err error)
	Update(ctx context.Context, e model.Employee) (model.Employee, error)
	Delete(ctx context.Context, id int) error
	Ping(ctx context.Context) error // used by /ready
	Close() error                   // called during shutdown
}

// MemoryEmployeeRepository keeps everything in memory (no database needed for this lesson).
type MemoryEmployeeRepository struct {
	mu     sync.RWMutex
	items  map[int]model.Employee
	nextID int
	closed bool
}

var _ EmployeeRepository = (*MemoryEmployeeRepository)(nil)

func NewMemoryEmployeeRepository() *MemoryEmployeeRepository {
	return &MemoryEmployeeRepository{items: map[int]model.Employee{}, nextID: 1}
}

// guard is called at the start of every operation (while holding the lock).
func (r *MemoryEmployeeRepository) guard(ctx context.Context) error {
	if err := ctx.Err(); err != nil { // cancelled or timed out
		return err
	}
	if r.closed {
		return ErrClosed
	}
	return nil
}

func (r *MemoryEmployeeRepository) emailTaken(email string, exceptID int) bool {
	for _, e := range r.items {
		if e.ID != exceptID && strings.EqualFold(e.Email, email) {
			return true
		}
	}
	return false
}

func (r *MemoryEmployeeRepository) Create(ctx context.Context, e model.Employee) (model.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.guard(ctx); err != nil {
		return model.Employee{}, err
	}
	if r.emailTaken(e.Email, 0) {
		return model.Employee{}, fmt.Errorf("%w: email %q is already in use", model.ErrConflict, e.Email)
	}
	e.ID = r.nextID
	e.CreatedAt = time.Now().UTC()
	r.nextID++
	r.items[e.ID] = e
	return e, nil
}

func (r *MemoryEmployeeRepository) GetByID(ctx context.Context, id int) (model.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if err := r.guard(ctx); err != nil {
		return model.Employee{}, err
	}
	e, ok := r.items[id]
	if !ok {
		return model.Employee{}, model.ErrNotFound
	}
	return e, nil
}

func (r *MemoryEmployeeRepository) List(ctx context.Context, f model.ListFilter) ([]model.Employee, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if err := r.guard(ctx); err != nil {
		return nil, 0, err
	}

	search := strings.ToLower(f.Search)
	matched := make([]model.Employee, 0, len(r.items))
	for _, e := range r.items {
		if search == "" ||
			strings.Contains(strings.ToLower(e.Name), search) ||
			strings.Contains(strings.ToLower(e.Email), search) ||
			strings.Contains(strings.ToLower(e.Department), search) {
			matched = append(matched, e)
		}
	}
	sort.Slice(matched, func(i, j int) bool { return matched[i].ID < matched[j].ID })

	total := len(matched)
	if f.Offset >= total {
		return []model.Employee{}, total, nil
	}
	matched = matched[f.Offset:]
	if f.Limit > 0 && f.Limit < len(matched) {
		matched = matched[:f.Limit]
	}
	return matched, total, nil
}

func (r *MemoryEmployeeRepository) Update(ctx context.Context, e model.Employee) (model.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.guard(ctx); err != nil {
		return model.Employee{}, err
	}
	old, ok := r.items[e.ID]
	if !ok {
		return model.Employee{}, model.ErrNotFound
	}
	if r.emailTaken(e.Email, e.ID) {
		return model.Employee{}, fmt.Errorf("%w: email %q is already in use", model.ErrConflict, e.Email)
	}
	e.CreatedAt = old.CreatedAt
	r.items[e.ID] = e
	return e, nil
}

func (r *MemoryEmployeeRepository) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.guard(ctx); err != nil {
		return err
	}
	if _, ok := r.items[id]; !ok {
		return model.ErrNotFound
	}
	delete(r.items, id)
	return nil
}

func (r *MemoryEmployeeRepository) Ping(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.guard(ctx)
}

// Close is part of the shutdown sequence. With a real database this would close the connection pool.
func (r *MemoryEmployeeRepository) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
	return nil
}
