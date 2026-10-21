package service

import (
	"errors"
	"math"
	"net/mail"
	"strings"

	"employee-management/model"
	"employee-management/repository"

	"gorm.io/gorm"
)

type EmployeeService interface {
	Create(*model.Employee) error
	GetAll(model.EmployeeFilter, model.EmployeeSort, int, int) (*model.EmployeeListResponse, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func validateEmployee(e *model.Employee) error {
	if e == nil {
		return errors.New("employee is required")
	}

	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)

	if e.Name == "" || len(e.Name) > 100 {
		return errors.New("name is required and must be at most 100 characters")
	}
	if len(e.Email) > 150 {
		return errors.New("email must be at most 150 characters")
	}
	addr, err := mail.ParseAddress(e.Email)
	if err != nil || addr.Address != e.Email {
		return errors.New("a valid email is required")
	}
	if e.Department == "" || len(e.Department) > 100 {
		return errors.New("department is required and must be at most 100 characters")
	}
	if math.IsNaN(e.Salary) || math.IsInf(e.Salary, 0) || e.Salary < 0 {
		return errors.New("salary must be a finite, non-negative number")
	}

	return nil
}

func (s *employeeService) Create(e *model.Employee) error {
	if err := validateEmployee(e); err != nil {
		return err
	}
	e.ID = 0 // Never accept a client-selected database ID.
	return s.repo.Create(e)
}

func (s *employeeService) GetAll(
	filter model.EmployeeFilter,
	sort model.EmployeeSort,
	page, pageSize int,
) (*model.EmployeeListResponse, error) {
	if page < 1 {
		return nil, errors.New("page must be at least 1")
	}
	if pageSize < 1 || pageSize > 100 {
		return nil, errors.New("pageSize must be between 1 and 100")
	}

	if sort.Field == "" {
		sort.Field = "id"
	}
	allowed := map[string]bool{
		"id": true, "name": true, "email": true,
		"department": true, "salary": true,
	}
	sort.Field = strings.ToLower(sort.Field)
	if !allowed[sort.Field] {
		return nil, errors.New("unsupported sort field")
	}
	if sort.Direction == "" {
		sort.Direction = "asc"
	}
	sort.Direction = strings.ToLower(sort.Direction)
	if sort.Direction != "asc" && sort.Direction != "desc" {
		return nil, errors.New("sort direction must be asc or desc")
	}

	employees, total, err := s.repo.GetAll(filter, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.EmployeeListResponse{
		Data: employees, Page: page, PageSize: pageSize,
		Total:      int(total),
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}, nil
}

func (s *employeeService) GetByID(id uint) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("employee ID must be positive")
	}
	return s.repo.GetByID(id)
}

func (s *employeeService) Update(id uint, e *model.Employee) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("employee ID must be positive")
	}
	if err := validateEmployee(e); err != nil {
		return nil, err
	}
	e.ID = id
	return s.repo.Update(id, e)
}

func (s *employeeService) Delete(id uint) error {
	if id == 0 {
		return errors.New("employee ID must be positive")
	}
	return s.repo.Delete(id)
}

// IsNotFound standardizes not-found checks for both API layers.
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound) ||
		(err != nil && err.Error() == "employee not found")
}

// IsDuplicate standardizes duplicate-email checks for both API layers.
func IsDuplicate(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}
