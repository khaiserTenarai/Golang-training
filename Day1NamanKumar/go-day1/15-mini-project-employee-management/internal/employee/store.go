package employee

import (
	"sort"
	"strings"
	"sync"
)

// Store is a concurrency-safe, in-memory employee repository.
type Store struct {
	mu     sync.RWMutex
	data   map[int]Employee
	nextID int
}

// NewStore returns an empty Store whose IDs start at 1.
func NewStore() *Store {
	return &Store{data: make(map[int]Employee), nextID: 1}
}

// Add validates e, assigns a new ID and stores it.
func (s *Store) Add(e Employee) (Employee, error) {
	e.Name = strings.TrimSpace(e.Name)
	e.Department = strings.TrimSpace(e.Department)
	if err := e.Validate(); err != nil {
		return Employee{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	e.ID = s.nextID
	s.data[e.ID] = e
	s.nextID++
	return e, nil
}

// GetByID returns the employee with the given ID, or ErrNotFound.
func (s *Store) GetByID(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.data[id]
	if !ok {
		return Employee{}, ErrNotFound
	}
	return e, nil
}

// SearchByName returns employees whose name contains query (case-insensitive).
func (s *Store) SearchByName(query string) []Employee {
	q := strings.ToLower(strings.TrimSpace(query))
	var out []Employee
	for _, e := range s.List() {
		if strings.Contains(strings.ToLower(e.Name), q) {
			out = append(out, e)
		}
	}
	return out
}

// List returns all employees sorted by ID.
func (s *Store) List() []Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Employee, 0, len(s.data))
	for _, e := range s.data {
		out = append(out, e)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Delete removes the employee with the given ID, or returns ErrNotFound.
func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return ErrNotFound
	}
	delete(s.data, id)
	return nil
}
