// Package service contains business logic.
package service

// Import context for request-scoped operations.
import "context"

// Import errors for validation.
import "errors"

// Import math for pagination calculations.
import "math"

// Import employee model.
import "employee-management/model"

// Import repository interface.
import "employee-management/repository"

type EmployeeService interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string, string, string, string, int, int) (*model.EmployeeListResponse, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
	Health(context.Context) error
}

type employeeService struct{ repository repository.EmployeeRepository }

// NewEmployeeService creates the service with dependency injection.
func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{repository: r}
}

// Create validates and creates an employee.
func (s *employeeService) Create(ctx context.Context, e *model.Employee) error {
	if err := validateEmployee(e); err != nil {
		return err
	}
	return s.repository.Create(ctx, e)
}

// GetAll applies filtering, sorting and pagination.
func (s *employeeService) GetAll(ctx context.Context, department, name, sortBy, sortOrder string, page, pageSize int) (*model.EmployeeListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	if sortOrder != "" && sortOrder != "asc" && sortOrder != "desc" {
		return nil, errors.New("sortOrder must be asc or desc")
	}
	items, total, err := s.repository.GetAll(ctx, department, name, sortBy, sortOrder, page, pageSize)
	if err != nil {
		return nil, err
	}
	return &model.EmployeeListResponse{Data: items, Page: page, PageSize: pageSize, Total: total, TotalPages: int(math.Ceil(float64(total) / float64(pageSize)))}, nil
}

// GetByID returns an employee.
func (s *employeeService) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than zero")
	}
	return s.repository.GetByID(ctx, id)
}

// Update validates and updates an employee.
func (s *employeeService) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("employee id must be greater than zero")
	}
	if err := validateEmployee(e); err != nil {
		return nil, err
	}
	return s.repository.Update(ctx, id, e)
}

// Delete deletes an employee.
func (s *employeeService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("employee id must be greater than zero")
	}
	return s.repository.Delete(ctx, id)
}

// Health checks the database.
func (s *employeeService) Health(ctx context.Context) error { return s.repository.Ping(ctx) }

// validateEmployee validates required employee fields.
func validateEmployee(e *model.Employee) error {
	if e == nil {
		return errors.New("employee body is required")
	}
	if e.Name == "" {
		return errors.New("employee name is required")
	}
	if e.Email == "" {
		return errors.New("employee email is required")
	}
	if e.Department == "" {
		return errors.New("employee department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
