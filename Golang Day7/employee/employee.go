package employee

import "errors"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

type Service struct {
	store map[int]Employee
}

func NewService() *Service {
	return &Service{
		store: make(map[int]Employee),
	}
}

func (s *Service) Add(emp Employee) error {
	if _, exists := s.store[emp.ID]; exists {
		return errors.New("employee already exists")
	}
	if emp.Name == "" {
		return errors.New("name cannot be empty")
	}
	s.store[emp.ID] = emp
	return nil
}

func (s *Service) Get(id int) (Employee, error) {
	emp, exists := s.store[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	return emp, nil
}
