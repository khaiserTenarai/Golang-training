package service

import (
	"employee-management/model"
	"employee-management/repository"
	"errors"
	"gorm.io/gorm"
	"net/mail"
	"strings"
	"unicode/utf8"
)

// EmployeeService is used by BOTH the REST controller and the GraphQL resolvers.
// All rules (validation, paging, sorting) live here, so the two APIs always behave the same.
type EmployeeService interface {
	Create(*model.Employee) error
	List(model.EmployeeFilter) (*model.EmployeePage, error)
	ListByDepartment(uint) ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}
type employeeService struct {
	repo  repository.EmployeeRepository
	depts repository.DepartmentRepository
}

func NewEmployeeService(r repository.EmployeeRepository, d repository.DepartmentRepository) EmployeeService {
	return &employeeService{r, d}
}

// validate checks the input of Create and Update.
func (s *employeeService) validate(e *model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	if e.Name == "" {
		return invalid("name is required")
	}
	if n := utf8.RuneCountInString(e.Name); n < 2 || n > 100 {
		return invalid("name must be 2 to 100 characters")
	}
	if a, err := mail.ParseAddress(e.Email); err != nil || a.Address != e.Email {
		return invalid("valid email is required")
	}
	if e.Salary < 0 {
		return invalid("salary cannot be negative")
	}
	if e.Salary > 10000000 {
		return invalid("salary is too large")
	}
	if e.DepartmentID == 0 {
		return invalid("departmentId is required")
	}
	if _, err := s.depts.GetByID(e.DepartmentID); err != nil {
		if IsNotFound(err) {
			return invalid("department does not exist")
		}
		return err
	}
	return nil
}

var validSortFields = map[string]bool{"id": true, "name": true, "email": true, "salary": true, "createdAt": true}

// normalizeFilter fills in defaults and rejects bad paging / sorting values.
func normalizeFilter(f *model.EmployeeFilter) error {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.PageSize == 0 {
		f.PageSize = 10
	}
	if f.Page < 1 {
		return invalid("page must be 1 or more")
	}
	if f.PageSize < 1 || f.PageSize > 100 {
		return invalid("pageSize must be between 1 and 100")
	}
	f.Search = strings.TrimSpace(f.Search)
	if f.SortBy == "" {
		f.SortBy = "id"
	}
	if !validSortFields[f.SortBy] {
		return invalid("sortBy must be one of: id, name, email, salary, createdAt")
	}
	f.SortOrder = strings.ToLower(f.SortOrder)
	if f.SortOrder == "" {
		f.SortOrder = "asc"
	}
	if f.SortOrder != "asc" && f.SortOrder != "desc" {
		return invalid("sortOrder must be asc or desc")
	}
	if f.MinSalary != nil && f.MaxSalary != nil && *f.MinSalary > *f.MaxSalary {
		return invalid("minSalary cannot be greater than maxSalary")
	}
	return nil
}
func saveError(err error) error {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrDuplicateEmail
	}
	return err
}
func (s *employeeService) Create(e *model.Employee) error {
	if err := s.validate(e); err != nil {
		return err
	}
	return saveError(s.repo.Create(e))
}

// List = pagination + filtering + sorting in one place.
func (s *employeeService) List(f model.EmployeeFilter) (*model.EmployeePage, error) {
	if err := normalizeFilter(&f); err != nil {
		return nil, err
	}
	items, total, err := s.repo.List(f)
	if err != nil {
		return nil, err
	}
	if items == nil {
		items = []model.Employee{}
	}
	// round up: 25 employees with pageSize 10 -> 3 pages
	pages := int((total + int64(f.PageSize) - 1) / int64(f.PageSize))
	return &model.EmployeePage{
		Items:      items,
		TotalCount: int(total),
		Page:       f.Page,
		PageSize:   f.PageSize,
		TotalPages: pages,
		HasNext:    f.Page < pages,
	}, nil
}
func (s *employeeService) ListByDepartment(deptID uint) ([]model.Employee, error) {
	return s.repo.ListByDepartment(deptID)
}
func (s *employeeService) GetByID(id uint) (*model.Employee, error) { return s.repo.GetByID(id) }
func (s *employeeService) Update(id uint, e *model.Employee) (*model.Employee, error) {
	if err := s.validate(e); err != nil {
		return nil, err
	}
	x, err := s.repo.Update(id, e)
	if err != nil {
		return nil, saveError(err)
	}
	return x, nil
}
func (s *employeeService) Delete(id uint) error { return s.repo.Delete(id) }
