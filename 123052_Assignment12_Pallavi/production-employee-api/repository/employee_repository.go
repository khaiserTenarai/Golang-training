package repository

import (
	"errors"
	"production-employee-api/models"
	"sort"
	"strings"
	"sync"
)

type EmployeeRepository interface {
	GetAll(params models.QueryParams) ([]models.Employee, int)
	GetByID(id string) (*models.Employee, error)
	Create(emp models.Employee) models.Employee
	IsReady() bool
}

type memoryRepo struct {
	mu        sync.RWMutex
	employees map[string]models.Employee
}

func NewEmployeeRepository() EmployeeRepository {
	repo := &memoryRepo{
		employees: make(map[string]models.Employee),
	}
	// Seed initial data
	repo.employees["1"] = models.Employee{ID: "1", Name: "Alice", Email: "alice@example.com", Department: "Engineering", Role: "admin", Salary: 90000}
	repo.employees["2"] = models.Employee{ID: "2", Name: "Bob", Email: "bob@example.com", Department: "HR", Role: "user", Salary: 60000}
	repo.employees["3"] = models.Employee{ID: "3", Name: "Charlie", Email: "charlie@example.com", Department: "Engineering", Role: "user", Salary: 75000}
	return repo
}

func (r *memoryRepo) IsReady() bool {
	return true
}

func (r *memoryRepo) GetAll(params models.QueryParams) ([]models.Employee, int) {
	r.RLock()
	defer r.RUnlock()

	var result []models.Employee
	for _, e := range r.employees {
		if params.Department != "" && !strings.EqualFold(e.Department, params.Department) {
			continue
		}
		result = append(result, e)
	}

	// Sorting
	if params.SortBy != "" {
		sort.Slice(result, func(i, j int) bool {
			if params.SortBy == "salary" {
				if params.Order == "desc" {
					return result[i].Salary > result[j].Salary
				}
				return result[i].Salary < result[j].Salary
			}
			if params.Order == "desc" {
				return result[i].Name > result[j].Name
			}
			return result[i].Name < result[j].Name
		})
	}

	total := len(result)

	// Pagination
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}

	start := (params.Page - 1) * params.Limit
	if start >= total {
		return []models.Employee{}, total
	}
	end := start + params.Limit
	if end > total {
		end = total
	}

	return result[start:end], total
}

func (r *memoryRepo) GetByID(id string) (*models.Employee, error) {
	r.RLock()
	defer r.RUnlock()
	emp, exists := r.employees[id]
	if !exists {
		return nil, errors.New("employee not found")
	}
	return &emp, nil
}

func (r *memoryRepo) RLock() {
	panic("unimplemented")
}

func (r *memoryRepo) RUnlock() {
	panic("unimplemented")
}

func (r *memoryRepo) Create(emp models.Employee) models.Employee {
	r.Lock()
	defer r.Unlock()
	r.employees[emp.ID] = emp
	return emp
}

func (r *memoryRepo) Lock() {
	panic("unimplemented")
}

func (r *memoryRepo) Unlock() {
	panic("unimplemented")
}
