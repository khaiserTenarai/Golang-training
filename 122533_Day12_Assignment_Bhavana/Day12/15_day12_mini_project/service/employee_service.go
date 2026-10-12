package service

import (
	"strings"

	"employeeapi/model"
	"employeeapi/repository"
	"employeeapi/validation"
)

// ValidationError carries every problem found, so the handler can
// show them all at once.
type ValidationError struct {
	Problems []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Problems, "; ")
}

type ListResult struct {
	Data  []model.Employee `json:"data"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
	Total int              `json:"total"`
}

// The service depends on the repository interface, not on Postgres.
type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func tidy(e model.Employee) model.Employee {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.ToLower(strings.TrimSpace(e.Email))
	e.Department = strings.TrimSpace(e.Department)
	return e
}

func (s *EmployeeService) Create(e model.Employee) (model.Employee, error) {
	e = tidy(e)
	if problems := validation.ValidateEmployee(e); len(problems) > 0 {
		return model.Employee{}, &ValidationError{Problems: problems}
	}
	return s.repo.Create(e)
}

func (s *EmployeeService) Get(id int) (model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) Update(id int, e model.Employee) (model.Employee, error) {
	e = tidy(e)
	if problems := validation.ValidateEmployee(e); len(problems) > 0 {
		return model.Employee{}, &ValidationError{Problems: problems}
	}
	e.ID = id
	if err := s.repo.Update(e); err != nil {
		return model.Employee{}, err
	}
	return e, nil
}

func (s *EmployeeService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *EmployeeService) List(f model.ListFilter) (ListResult, error) {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.Limit < 1 {
		f.Limit = 10
	}
	if f.Limit > 100 {
		f.Limit = 100
	}

	employees, total, err := s.repo.List(f)
	if err != nil {
		return ListResult{}, err
	}
	return ListResult{Data: employees, Page: f.Page, Limit: f.Limit, Total: total}, nil
}

func (s *EmployeeService) Ready() error {
	return s.repo.Ping()
}
