package service

import (
	"fmt"
	"sync"

	"graphql_employee/models"
)

type DepartmentService struct {
	mu   sync.RWMutex
	data map[int]models.Department
	next int
}

func NewDepartmentService() *DepartmentService {
	s := &DepartmentService{data: make(map[int]models.Department), next: 1}
	seed := []models.Department{
		{Name: "Engineering", Location: "Building A"},
		{Name: "Design", Location: "Building B"},
		{Name: "Marketing", Location: "Building C"},
		{Name: "HR", Location: "Building D"},
		{Name: "Finance", Location: "Building A"},
	}
	for _, d := range seed {
		d.ID = s.next
		s.data[d.ID] = d
		s.next++
	}
	return s
}

func (s *DepartmentService) GetByID(id int) (models.Department, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d, ok := s.data[id]
	if !ok {
		return models.Department{}, fmt.Errorf("department %d not found", id)
	}
	return d, nil
}

func (s *DepartmentService) GetAll() []models.Department {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var list []models.Department
	for _, d := range s.data {
		list = append(list, d)
	}
	return list
}

// Task 5: Batch load departments for optimization
func (s *DepartmentService) GetByIDs(ids []int) map[int]models.Department {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make(map[int]models.Department)
	for _, id := range ids {
		if d, ok := s.data[id]; ok {
			result[id] = d
		}
	}
	return result
}
