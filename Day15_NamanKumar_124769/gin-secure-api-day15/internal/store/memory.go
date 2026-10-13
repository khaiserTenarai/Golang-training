package store

import (
	"errors"
	"sort"
	"strings"
	"sync"

	"gin-secure-api-day15/internal/model"
)

var ErrNotFound = errors.New("employee not found")
var ErrConflict = errors.New("email already exists")

type EmployeeStore interface {
	List(search, department string, offset, limit int) ([]model.Employee, int)
	Get(id int) (model.Employee, bool)
	Create(employee model.Employee) (model.Employee, error)
	Update(id int, employee model.Employee) (model.Employee, error)
	Delete(id int) bool
}

type MemoryStore struct {
	mu        sync.RWMutex
	nextID    int
	employees map[int]model.Employee
}

func NewMemoryStore() *MemoryStore {
	s := &MemoryStore{nextID: 3, employees: map[int]model.Employee{}}
	s.employees[1] = model.Employee{ID: 1, Name: "Asha Sharma", Email: "asha@example.com", Department: "Engineering", Role: "employee"}
	s.employees[2] = model.Employee{ID: 2, Name: "Ravi Singh", Email: "ravi@example.com", Department: "Operations", Role: "manager"}
	return s
}

func (s *MemoryStore) List(search, department string, offset, limit int) ([]model.Employee, int) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	all := make([]model.Employee, 0, len(s.employees))
	search = strings.ToLower(strings.TrimSpace(search))
	department = strings.ToLower(strings.TrimSpace(department))
	for _, employee := range s.employees {
		if search != "" && !strings.Contains(strings.ToLower(employee.Name+" "+employee.Email), search) {
			continue
		}
		if department != "" && strings.ToLower(employee.Department) != department {
			continue
		}
		all = append(all, employee)
	}
	sort.Slice(all, func(i, j int) bool { return all[i].ID < all[j].ID })
	total := len(all)
	if offset >= total {
		return []model.Employee{}, total
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return all[offset:end], total
}

func (s *MemoryStore) Get(id int) (model.Employee, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	employee, ok := s.employees[id]
	return employee, ok
}

func (s *MemoryStore) Create(employee model.Employee) (model.Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, existing := range s.employees {
		if strings.EqualFold(existing.Email, employee.Email) {
			return model.Employee{}, ErrConflict
		}
	}
	employee.ID = s.nextID
	s.nextID++
	s.employees[employee.ID] = employee
	return employee, nil
}

func (s *MemoryStore) Update(id int, employee model.Employee) (model.Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.employees[id]; !ok {
		return model.Employee{}, ErrNotFound
	}
	for existingID, existing := range s.employees {
		if existingID != id && strings.EqualFold(existing.Email, employee.Email) {
			return model.Employee{}, ErrConflict
		}
	}
	employee.ID = id
	s.employees[id] = employee
	return employee, nil
}

func (s *MemoryStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.employees[id]; !ok {
		return false
	}
	delete(s.employees, id)
	return true
}
