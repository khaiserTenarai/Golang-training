package service

import (
	"context"
	"day13_project/config"
	"day13_project/model"
	"day13_project/repository"
	"fmt"
	"sort"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
	cfg  *config.Config
}

func NewEmployeeService(repo *repository.EmployeeRepository, cfg *config.Config) *EmployeeService {
	return &EmployeeService{repo: repo, cfg: cfg}
}

func (s *EmployeeService) ListEmployees(ctx context.Context, sortAlpha bool, page, pageSize int) ([]model.Employee, error) {
	list, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if sortAlpha {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name < list[j].Name
		})
	}

	// Pagination math safety rules
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	start := (page - 1) * pageSize
	if start > len(list) {
		return []model.Employee{}, nil
	}
	end := start + pageSize
	if end > len(list) {
		end = len(list)
	}

	return list[start:end], nil
}

func (s *EmployeeService) AddEmployee(ctx context.Context, emp model.Employee) error {
	if s.cfg.EnableNewValidation {
		if err := emp.Validate(); err != nil {
			return fmt.Errorf("validation error layer exception: %w", err)
		}
	}
	return s.repo.Save(ctx, emp)
}
