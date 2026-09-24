package employee

import "errors"

// ErrNotFound is returned when an employee with the given ID does not exist.
var ErrNotFound = errors.New("employee not found")

// Employee holds the details of a single employee.
type Employee struct {
	// ID is the unique, auto-assigned identifier.
	ID int
	// Name is the employee's full name.
	Name string
	// Department is the team the employee belongs to.
	Department string
	// Salary is the annual salary in INR.
	Salary float64
}

// Store keeps employees in memory, keyed by ID.
// The zero value is not ready for use; create one with [NewStore].
type Store struct {
	data   map[int]Employee
	nextID int
}

// NewStore returns an empty, ready-to-use Store.
func NewStore() *Store {
	return &Store{data: make(map[int]Employee), nextID: 1}
}

// Add creates a new employee, assigns it the next ID and returns it.
func (s *Store) Add(name, dept string, salary float64) Employee {
	e := Employee{ID: s.nextID, Name: name, Department: dept, Salary: salary}
	s.data[e.ID] = e
	s.nextID++
	return e
}

// Get returns the employee with the given id, or [ErrNotFound].
func (s *Store) Get(id int) (Employee, error) {
	e, ok := s.data[id]
	if !ok {
		return Employee{}, ErrNotFound
	}
	return e, nil
}

// Count reports how many employees are in the store.
func (s *Store) Count() int { return len(s.data) }
