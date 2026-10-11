package repository

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"configurable_api/models"
)

type EmployeeRepository interface {
	Create(ctx context.Context, emp models.Employee) (models.Employee, error)
	GetAll(ctx context.Context, params models.QueryParams) ([]models.Employee, int, error)
	GetByID(ctx context.Context, id int) (models.Employee, error)
	Update(ctx context.Context, id int, emp models.Employee) (models.Employee, error)
	Delete(ctx context.Context, id int) error
	Count() int
}

type InMemoryRepo struct {
	mu   sync.RWMutex
	data map[int]models.Employee
	next int
}

func NewInMemoryRepo() EmployeeRepository {
	r := &InMemoryRepo{data: make(map[int]models.Employee), next: 1}
	seed := []models.Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, Department: "Engineering", Salary: 60000},
		{Name: "Alice", Email: "alice@test.com", Age: 30, Department: "Design", Salary: 55000},
		{Name: "Bob", Email: "bob@test.com", Age: 28, Department: "Marketing", Salary: 45000},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, Department: "Engineering", Salary: 80000},
	}
	for _, e := range seed {
		e.ID = r.next
		e.CreatedAt = time.Now()
		r.data[e.ID] = e
		r.next++
	}
	return r
}

// Task 11: Context passed through to repository
func (r *InMemoryRepo) Create(ctx context.Context, emp models.Employee) (models.Employee, error) {
	select {
	case <-ctx.Done():
		return models.Employee{}, ctx.Err()
	default:
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, e := range r.data {
		if strings.EqualFold(e.Email, emp.Email) {
			return models.Employee{}, fmt.Errorf("email already exists")
		}
	}
	emp.ID = r.next
	r.next++
	emp.CreatedAt = time.Now()
	r.data[emp.ID] = emp
	return emp, nil
}

func (r *InMemoryRepo) GetAll(ctx context.Context, p models.QueryParams) ([]models.Employee, int, error) {
	select {
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	default:
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []models.Employee
	for _, e := range r.data {
		if p.Name != "" && !strings.Contains(strings.ToLower(e.Name), strings.ToLower(p.Name)) {
			continue
		}
		if p.Department != "" && !strings.EqualFold(e.Department, p.Department) {
			continue
		}
		if p.MinSalary > 0 && e.Salary < p.MinSalary {
			continue
		}
		list = append(list, e)
	}

	desc := strings.ToLower(p.SortOrder) == "desc"
	sort.Slice(list, func(i, j int) bool {
		var less bool
		switch strings.ToLower(p.SortBy) {
		case "name":
			less = list[i].Name < list[j].Name
		case "salary":
			less = list[i].Salary < list[j].Salary
		case "age":
			less = list[i].Age < list[j].Age
		default:
			less = list[i].ID < list[j].ID
		}
		if desc {
			return !less
		}
		return less
	})

	total := len(list)
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	s := (p.Page - 1) * p.PageSize
	e := s + p.PageSize
	if s > total {
		s = total
	}
	if e > total {
		e = total
	}
	return list[s:e], total, nil
}

func (r *InMemoryRepo) GetByID(ctx context.Context, id int) (models.Employee, error) {
	select {
	case <-ctx.Done():
		return models.Employee{}, ctx.Err()
	default:
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	emp, ok := r.data[id]
	if !ok {
		return models.Employee{}, fmt.Errorf("not found")
	}
	return emp, nil
}

func (r *InMemoryRepo) Update(ctx context.Context, id int, emp models.Employee) (models.Employee, error) {
	select {
	case <-ctx.Done():
		return models.Employee{}, ctx.Err()
	default:
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	old, ok := r.data[id]
	if !ok {
		return models.Employee{}, fmt.Errorf("not found")
	}
	emp.ID = id
	emp.CreatedAt = old.CreatedAt
	r.data[id] = emp
	return emp, nil
}

func (r *InMemoryRepo) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return fmt.Errorf("not found")
	}
	delete(r.data, id)
	return nil
}

func (r *InMemoryRepo) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.data)
}
