package service

import (
	"context"
	"errors"
	"math"
	"net/mail"
	"strings"

	"employee-management/model"
	"employee-management/repository"
)

type EmployeeService interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string, string, string, string, int, int) (*model.EmployeeListResponse, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
	Health(context.Context) error
}

type employeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{repository: r}
}

func (s *employeeService) Create(
	ctx context.Context,
	e *model.Employee,
) error {
	if err := validateEmployee(e); err != nil {
		return err
	}

	return s.repository.Create(ctx, e)
}

func (s *employeeService) GetAll(
	ctx context.Context,
	department, name, sortBy, sortOrder string,
	page, pageSize int,
) (*model.EmployeeListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	if sortOrder != "" &&
		!strings.EqualFold(sortOrder, "asc") &&
		!strings.EqualFold(sortOrder, "desc") {
		return nil, errors.New("sort_order must be asc or desc")
	}

	employees, total, err := s.repository.GetAll(
		ctx, department, name, sortBy, sortOrder, page, pageSize,
	)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &model.EmployeeListResponse{
		Data:       employees,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *employeeService) GetByID(
	ctx context.Context,
	id int64,
) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee ID must be greater than zero")
	}

	return s.repository.GetByID(ctx, id)
}

func (s *employeeService) Update(
	ctx context.Context,
	id int64,
	e *model.Employee,
) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee ID must be greater than zero")
	}
	if err := validateEmployee(e); err != nil {
		return nil, err
	}

	return s.repository.Update(ctx, id, e)
}

func (s *employeeService) Delete(
	ctx context.Context,
	id int64,
) error {
	if id <= 0 {
		return errors.New("employee ID must be greater than zero")
	}

	return s.repository.Delete(ctx, id)
}

func (s *employeeService) Health(ctx context.Context) error {
	return s.repository.Ping(ctx)
}

func validateEmployee(e *model.Employee) error {
	if e == nil {
		return errors.New("employee body is required")
	}
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("employee name is required")
	}
	if _, err := mail.ParseAddress(e.Email); err != nil {
		return errors.New("valid employee email is required")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("employee department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return nil
}
