// Package store is the private in-memory employee repository.
package store

// Employee is a stored employee record.
type Employee struct {
	ID    int
	Name  string
	Email string
}

// Store holds employees in memory.
type Store struct {
	items  []Employee
	nextID int
}

// New returns an empty Store.
func New() *Store { return &Store{nextID: 1} }

// Add inserts an employee and returns it with its assigned ID.
func (s *Store) Add(name, email string) Employee {
	e := Employee{ID: s.nextID, Name: name, Email: email}
	s.items = append(s.items, e)
	s.nextID++
	return e
}
