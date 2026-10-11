// Task 3: Repository Layer — data access abstraction.
// Task 4: Dependency injection via interface.
// The repository exposes an interface so the service layer
// depends on abstraction, not a concrete implementation.

package repository

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"production_api/models"
)

// ---- Repository Interface (Task 4: DI contract) ----

type EmployeeRepository interface {
	Create(emp models.Employee) (models.Employee, error)
	GetAll(params models.QueryParams) ([]models.Employee, int, error)
	GetByID(id int) (models.Employee, error)
	Update(id int, emp models.Employee) (models.Employee, error)
	Delete(id int) error
	Count() int
}

// ---- In-Memory Implementation ----

type InMemoryEmployeeRepo struct {
	mu        sync.RWMutex
	employees map[int]models.Employee
	nextID    int
}

func NewInMemoryEmployeeRepo() EmployeeRepository {
	repo := &InMemoryEmployeeRepo{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}
	// Seed data
	seed := []models.Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, Department: "Engineering", Salary: 60000, Role: "developer"},
		{Name: "Alice", Email: "alice@test.com", Age: 30, Department: "Design", Salary: 55000, Role: "designer"},
		{Name: "Bob", Email: "bob@test.com", Age: 28, Department: "Marketing", Salary: 45000, Role: "marketer"},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, Department: "Engineering", Salary: 80000, Role: "lead"},
		{Name: "Diana", Email: "diana@test.com", Age: 27, Department: "HR", Salary: 50000, Role: "hr"},
	}
	for _, e := range seed {
		e.ID = repo.nextID
		e.CreatedAt = time.Now()
		e.UpdatedAt = time.Now()
		repo.employees[e.ID] = e
		repo.nextID++
	}
	return repo
}

func (r *InMemoryEmployeeRepo) Create(emp models.Employee) (models.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check duplicate email
	for _, e := range r.employees {
		if strings.EqualFold(e.Email, emp.Email) {
			return models.Employee{}, fmt.Errorf("email '%s' already exists", emp.Email)
		}
	}

	emp.ID = r.nextID
	r.nextID++
	emp.CreatedAt = time.Now()
	emp.UpdatedAt = time.Now()
	r.employees[emp.ID] = emp
	return emp, nil
}

// GetAll with filtering (Task 11), sorting (Task 12), pagination (Task 10)
func (r *InMemoryEmployeeRepo) GetAll(params models.QueryParams) ([]models.Employee, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Task 11: Filtering
	var filtered []models.Employee
	for _, emp := range r.employees {
		if params.Name != "" && !strings.Contains(strings.ToLower(emp.Name), strings.ToLower(params.Name)) {
			continue
		}
		if params.Department != "" && !strings.EqualFold(emp.Department, params.Department) {
			continue
		}
		if params.MinSalary > 0 && emp.Salary < params.MinSalary {
			continue
		}
		if params.MaxSalary > 0 && emp.Salary > params.MaxSalary {
			continue
		}
		filtered = append(filtered, emp)
	}

	// Task 12: Sorting
	sortBy := strings.ToLower(params.SortBy)
	desc := strings.ToLower(params.SortOrder) == "desc"
	sort.Slice(filtered, func(i, j int) bool {
		var less bool
		switch sortBy {
		case "name":
			less = strings.ToLower(filtered[i].Name) < strings.ToLower(filtered[j].Name)
		case "salary":
			less = filtered[i].Salary < filtered[j].Salary
		case "age":
			less = filtered[i].Age < filtered[j].Age
		case "department":
			less = strings.ToLower(filtered[i].Department) < strings.ToLower(filtered[j].Department)
		default:
			less = filtered[i].ID < filtered[j].ID
		}
		if desc {
			return !less
		}
		return less
	})

	// Task 10: Pagination
	total := len(filtered)
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 {
		params.PageSize = 10
	}
	start := (params.Page - 1) * params.PageSize
	end := start + params.PageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	return filtered[start:end], total, nil
}

func (r *InMemoryEmployeeRepo) GetByID(id int) (models.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	emp, ok := r.employees[id]
	if !ok {
		return models.Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return emp, nil
}

func (r *InMemoryEmployeeRepo) Update(id int, emp models.Employee) (models.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, ok := r.employees[id]
	if !ok {
		return models.Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	// Check duplicate email (skip self)
	for _, e := range r.employees {
		if e.ID != id && strings.EqualFold(e.Email, emp.Email) {
			return models.Employee{}, fmt.Errorf("email '%s' already exists", emp.Email)
		}
	}
	emp.ID = id
	emp.CreatedAt = existing.CreatedAt
	emp.UpdatedAt = time.Now()
	r.employees[id] = emp
	return emp, nil
}

func (r *InMemoryEmployeeRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.employees[id]; !ok {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	delete(r.employees, id)
	return nil
}

func (r *InMemoryEmployeeRepo) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.employees)
}

// Helper used by handler for total pages
func TotalPages(total, pageSize int) int {
	return int(math.Ceil(float64(total) / float64(pageSize)))
}
