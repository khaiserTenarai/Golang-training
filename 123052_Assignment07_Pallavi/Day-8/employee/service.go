package employee

import (
	"errors"
	"log/slog"
	"sync"
)

var (
	ErrDuplicate = errors.New("employee already exists")
	ErrNotFound  = errors.New("employee not found")
)

type Service struct {
	mu   sync.RWMutex
	data map[int]Employee
}

func NewService() *Service {
	return &Service{data: make(map[int]Employee)}
}

func (s *Service) Add(e Employee) error {
	if err := e.Validate(); err != nil {
		slog.Warn("validation failed", "id", e.ID, "err", err)
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[e.ID]; ok {
		return ErrDuplicate
	}
	s.data[e.ID] = e
	slog.Info("employee added", "id", e.ID)
	return nil
}

func (s *Service) Get(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.data[id]
	if !ok {
		return Employee{}, ErrNotFound
	}
	return e, nil
}

func (s *Service) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}