package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

var (
	ErrNotFound       = errors.New("employee not found")
	ErrDuplicateEmail = errors.New("email already exists")
)

type Employee struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Dept   string
	Salary float64
}

var validDepts = map[string]bool{"Eng": true, "HR": true, "Ops": true, "Sales": true}

func (e Employee) Validate() error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(e.Email, "@") || !strings.Contains(e.Email, ".") {
		return fmt.Errorf("invalid email %q", e.Email)
	}
	if e.Age < 18 || e.Age > 65 {
		return fmt.Errorf("age %d out of range (18-65)", e.Age)
	}
	if !validDepts[e.Dept] {
		return fmt.Errorf("invalid department %q", e.Dept)
	}
	if e.Salary <= 0 {
		return errors.New("salary must be positive")
	}
	return nil
}

func CalculateTax(annual float64) float64 {
	switch {
	case annual <= 250000:
		return 0
	case annual <= 500000:
		return (annual - 250000) * 0.05
	default:
		return 12500 + (annual-500000)*0.20
	}
}

func CalculateSalary(base, bonusPct float64) (gross, tax, net float64, err error) {
	if base <= 0 {
		return 0, 0, 0, errors.New("base salary must be positive")
	}
	if bonusPct < 0 {
		return 0, 0, 0, errors.New("bonus percent cannot be negative")
	}
	gross = base + base*bonusPct/100
	tax = CalculateTax(gross)
	net = gross - tax
	return gross, tax, net, nil
}

type Service struct {
	mu     sync.RWMutex
	data   map[int]Employee
	nextID int
}

func NewService() *Service {
	return &Service{data: make(map[int]Employee), nextID: 1}
}

func (s *Service) Add(e Employee) (Employee, error) {
	if err := e.Validate(); err != nil {
		logger.Warn("validation failed", "name", e.Name, "err", err)
		return Employee{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, existing := range s.data {
		if strings.EqualFold(existing.Email, e.Email) {
			logger.Warn("duplicate email", "email", e.Email)
			return Employee{}, ErrDuplicateEmail
		}
	}

	e.ID = s.nextID
	s.nextID++
	s.data[e.ID] = e
	logger.Info("employee added", "id", e.ID, "name", e.Name)
	return e, nil
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

func (s *Service) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return ErrNotFound
	}
	delete(s.data, id)
	logger.Info("employee deleted", "id", id)
	return nil
}

func (s *Service) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}
