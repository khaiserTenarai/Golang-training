package service

import (
	"context"
	"employee_api/config"
	"employee_api/model"
	"employee_api/repository"
	"errors"
	"sort"
	"strings"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
	cfg  *config.Config
}

func NewEmployeeService(repo repository.EmployeeRepository, cfg *config.Config) *EmployeeService {
	return &EmployeeService{repo: repo, cfg: cfg}
}

func (s *EmployeeService) AddEmployee(ctx context.Context, emp model.Employee) (model.Employee, error) {
	if s.cfg.EnableNewValidation {
		if err := emp.Validate(); err != nil {
			return model.Employee{}, err
		}
	}
	return s.repo.Insert(ctx, emp), nil
}

func (s *EmployeeService) SearchAndList(ctx context.Context, deptFilter string, page, pageSize int) ([]model.Employee, error) {
	all := s.repo.FindAll(ctx)

	// Filter Logic
	var filtered []model.Employee
	for _, e := range all {
		if deptFilter == "" || strings.EqualFold(e.Department, deptFilter) {
			filtered = append(filtered, e)
		}
	}

	// Sort explicitly by Name alphabetically
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Pagination Math
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	start := (page - 1) * pageSize
	if start >= len(filtered) {
		return []model.Employee{}, nil
	}
	end := start + pageSize
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}
