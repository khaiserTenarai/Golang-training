// Package service contains business logic.
package service

// Import context for request-scoped operations.
import "context"

// Import errors for validation and business errors.
import "errors"

// Import employee model.
import "employee-management/model"

// Import repository interface.
import "employee-management/repository"

// EmployeeService defines employee business operations.
type EmployeeService interface {
	// Create validates and creates an employee.
	Create(ctx context.Context, employee *model.Employee) error
	// GetAll returns all employees.
	GetAll(ctx context.Context, sortBy string) ([]model.Employee, error)
	// GetByID returns an employee by ID.
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	// Update validates and updates an employee.
	Update(ctx context.Context, id int64, employee *model.Employee) (*model.Employee, error)
	// Delete deletes an employee by ID.
	Delete(ctx context.Context, id int64) error
}

// employeeService implements EmployeeService.
type employeeService struct {
	// repository stores the repository dependency.
	repository repository.EmployeeRepository
}

// NewEmployeeService creates the service with dependency injection.
func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	// Return the service implementation.
	return &employeeService{repository: repository}
}

// Create validates and creates an employee.
func (s *employeeService) Create(ctx context.Context, employee *model.Employee) error {
	// Validate the employee before saving.
	if err := validateEmployee(employee); err != nil {
		// Return validation error to the controller.
		return err
	}
	// Delegate persistence to the repository.
	return s.repository.Create(ctx, employee)
}

// GetAll retrieves all employees.
func (s *employeeService) GetAll(ctx context.Context, sortBy string) ([]model.Employee, error) {
	// Delegate retrieval to the repository.
	return s.repository.GetAll(ctx, sortBy)
}

// GetByID retrieves one employee.
func (s *employeeService) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	// Validate the path parameter.
	if id <= 0 {
		// Return an error for invalid IDs.
		return nil, errors.New("employee id must be greater than zero")
	}
	// Delegate retrieval to the repository.
	return s.repository.GetByID(ctx, id)
}

// Update validates and updates an employee.
func (s *employeeService) Update(ctx context.Context, id int64, employee *model.Employee) (*model.Employee, error) {
	// Validate the path parameter.
	if id <= 0 {
		// Return an error for invalid IDs.
		return nil, errors.New("employee id must be greater than zero")
	}
	// Validate the request body.
	if err := validateEmployee(employee); err != nil {
		// Return validation error.
		return nil, err
	}
	// Delegate update to the repository.
	return s.repository.Update(ctx, id, employee)
}

// Delete deletes an employee.
func (s *employeeService) Delete(ctx context.Context, id int64) error {
	// Validate the path parameter.
	if id <= 0 {
		// Return an error for invalid IDs.
		return errors.New("employee id must be greater than zero")
	}
	// Delegate deletion to the repository.
	return s.repository.Delete(ctx, id)
}

// validateEmployee validates fields required by the business layer.
func validateEmployee(employee *model.Employee) error {
	// Check for a missing employee object.
	if employee == nil {
		// Return an invalid request error.
		return errors.New("employee body is required")
	}
	// Check the employee name.
	if employee.Name == "" {
		// Return a name validation error.
		return errors.New("employee name is required")
	}
	// Check the employee email.
	if employee.Email == "" {
		// Return an email validation error.
		return errors.New("employee email is required")
	}
	// Check the employee department.
	if employee.Department == "" {
		// Return a department validation error.
		return errors.New("employee department is required")
	}
	// Check the salary.
	if employee.Salary < 0 {
		// Return a salary validation error.
		return errors.New("salary cannot be negative")
	}
	// Return nil when validation succeeds.
	return nil
}
