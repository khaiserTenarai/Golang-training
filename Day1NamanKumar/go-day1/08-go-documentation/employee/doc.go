// Package employee provides an in-memory store for managing employees.
//
// A Store is safe to use from a single goroutine. Typical usage:
//
//	s := employee.NewStore()
//	e := s.Add("Asha Rao", "Engineering", 85000)
//	found, err := s.Get(e.ID)
//
// Errors returned by the store can be compared with [errors.Is] against
// [ErrNotFound].
package employee
