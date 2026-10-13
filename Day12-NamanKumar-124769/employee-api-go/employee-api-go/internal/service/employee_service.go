package service

import (
	"errors"
	"math"
	"sort"
	"strings"

	"employee-api-go/internal/domain"
	"employee-api-go/internal/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(input domain.CreateEmployeeInput) (domain.Employee, error) {
	return s.repo.Create(input)
}

func (s *EmployeeService) GetByID(id int64) (domain.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) Update(id int64, input domain.UpdateEmployeeInput) (domain.Employee, error) {
	return s.repo.Update(id, input)
}

func (s *EmployeeService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *EmployeeService) List(query domain.EmployeeQuery) (domain.EmployeePage, error) {
	if query.Page < 1 {
		return domain.EmployeePage{}, errors.New("page must be greater than zero")
	}
	if query.Limit < 1 || query.Limit > 100 {
		return domain.EmployeePage{}, errors.New("limit must be between 1 and 100")
	}

	if query.SortBy == "" {
		query.SortBy = "id"
	}
	if query.Order == "" {
		query.Order = "asc"
	}
	if !allowedSortFields[query.SortBy] {
		return domain.EmployeePage{}, errors.New("sort_by must be one of id, name, email, department, created_at")
	}
	if query.Order != "asc" && query.Order != "desc" {
		return domain.EmployeePage{}, errors.New("order must be asc or desc")
	}

	employees := s.repo.List()
	filtered := make([]domain.Employee, 0, len(employees))
	for _, employee := range employees {
		if query.Department != "" && !strings.EqualFold(employee.Department, query.Department) {
			continue
		}
		if query.Search != "" {
			search := strings.ToLower(query.Search)
			if !strings.Contains(strings.ToLower(employee.Name), search) && !strings.Contains(strings.ToLower(employee.Email), search) {
				continue
			}
		}
		filtered = append(filtered, employee)
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		left, right := filtered[i], filtered[j]
		var less bool
		switch query.SortBy {
		case "name":
			less = strings.ToLower(left.Name) < strings.ToLower(right.Name)
		case "email":
			less = strings.ToLower(left.Email) < strings.ToLower(right.Email)
		case "department":
			less = strings.ToLower(left.Department) < strings.ToLower(right.Department)
		case "created_at":
			less = left.CreatedAt.Before(right.CreatedAt)
		default:
			less = left.ID < right.ID
		}
		if query.Order == "desc" {
			return !less && !employeeSortEqual(left, right, query.SortBy)
		}
		return less
	})

	total := len(filtered)
	start := (query.Page - 1) * query.Limit
	if start > total {
		start = total
	}
	end := start + query.Limit
	if end > total {
		end = total
	}
	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	if totalPages == 0 {
		totalPages = 0
	}

	return domain.EmployeePage{
		Data:       filtered[start:end],
		Page:       query.Page,
		Limit:      query.Limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

var allowedSortFields = map[string]bool{
	"id": true, "name": true, "email": true, "department": true, "created_at": true,
}

func employeeSortEqual(a, b domain.Employee, field string) bool {
	switch field {
	case "name":
		return strings.EqualFold(a.Name, b.Name)
	case "email":
		return strings.EqualFold(a.Email, b.Email)
	case "department":
		return strings.EqualFold(a.Department, b.Department)
	case "created_at":
		return a.CreatedAt.Equal(b.CreatedAt)
	default:
		return a.ID == b.ID
	}
}

var _ = repository.ErrNotFound
