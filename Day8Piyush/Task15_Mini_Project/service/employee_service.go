// Package service implements the Employee Service with validation,
// CRUD operations, and salary calculations.
// This service is race-safe using sync.RWMutex.
package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"employee_service/models"
)

// EmployeeService manages employees with thread-safe operations.
type EmployeeService struct {
	mu        sync.RWMutex
	employees map[int]models.Employee
	nextID    int
}

// NewEmployeeService creates a new service instance.
func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}
}

// ---------- Validation ----------

// ValidateEmployee checks all business rules for an employee.
func ValidateEmployee(emp models.Employee) error {
	if strings.TrimSpace(emp.Name) == "" {
		return errors.New("name is required")
	}
	if len(emp.Name) < 2 {
		return errors.New("name must be at least 2 characters")
	}
	if strings.TrimSpace(emp.Email) == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(emp.Email) {
		return errors.New("invalid email format")
	}
	if emp.Age < 18 {
		return errors.New("age must be at least 18")
	}
	if emp.Age > 65 {
		return errors.New("age must be at most 65")
	}
	if strings.TrimSpace(emp.Department) == "" {
		return errors.New("department is required")
	}
	if emp.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}

func isValidEmail(email string) bool {
	atIdx := strings.Index(email, "@")
	if atIdx <= 0 {
		return false
	}
	domain := email[atIdx+1:]
	dotIdx := strings.Index(domain, ".")
	return dotIdx > 0 && dotIdx < len(domain)-1
}

// ---------- CRUD ----------

// Add validates and adds a new employee. Returns the created employee.
func (s *EmployeeService) Add(emp models.Employee) (models.Employee, error) {
	if err := ValidateEmployee(emp); err != nil {
		return models.Employee{}, fmt.Errorf("validation failed: %w", err)
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check duplicate email
	for _, existing := range s.employees {
		if strings.EqualFold(existing.Email, emp.Email) {
			return models.Employee{}, errors.New("email already exists")
		}
	}

	emp.ID = s.nextID
	s.nextID++
	s.employees[emp.ID] = emp
	return emp, nil
}

// Get retrieves an employee by ID.
func (s *EmployeeService) Get(id int) (models.Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	emp, exists := s.employees[id]
	if !exists {
		return models.Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return emp, nil
}

// GetAll returns all employees.
func (s *EmployeeService) GetAll() []models.Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		result = append(result, emp)
	}
	return result
}

// Update replaces an employee's data after validation.
func (s *EmployeeService) Update(id int, emp models.Employee) (models.Employee, error) {
	if err := ValidateEmployee(emp); err != nil {
		return models.Employee{}, fmt.Errorf("validation failed: %w", err)
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.employees[id]; !exists {
		return models.Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}

	// Check duplicate email (skip self)
	for _, existing := range s.employees {
		if existing.ID != id && strings.EqualFold(existing.Email, emp.Email) {
			return models.Employee{}, errors.New("email already exists")
		}
	}

	emp.ID = id
	s.employees[id] = emp
	return emp, nil
}

// Delete removes an employee by ID.
func (s *EmployeeService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.employees[id]; !exists {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	delete(s.employees, id)
	return nil
}

// Count returns the number of employees.
func (s *EmployeeService) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.employees)
}

// ---------- Salary Operations ----------

// CalculateNetSalary computes net salary after tax deduction.
func CalculateNetSalary(gross float64, taxPercent float64) (float64, error) {
	if gross < 0 {
		return 0, errors.New("gross salary cannot be negative")
	}
	if taxPercent < 0 || taxPercent > 100 {
		return 0, errors.New("tax percent must be between 0 and 100")
	}
	return gross - (gross * taxPercent / 100), nil
}

// CalculateAnnualSalary returns monthly * 12.
func CalculateAnnualSalary(monthly float64) float64 {
	return monthly * 12
}

// CalculateBonus returns bonus based on performance rating (1-5).
func CalculateBonus(salary float64, rating int) (float64, error) {
	if salary < 0 {
		return 0, errors.New("salary cannot be negative")
	}
	if rating < 1 || rating > 5 {
		return 0, errors.New("rating must be between 1 and 5")
	}
	rates := [6]float64{0, 0, 5, 10, 15, 20}
	return salary * rates[rating] / 100, nil
}

// GetTaxBracket returns the tax bracket label for a salary.
func GetTaxBracket(salary float64) string {
	switch {
	case salary <= 250000:
		return "No Tax"
	case salary <= 500000:
		return "5%"
	case salary <= 1000000:
		return "20%"
	default:
		return "30%"
	}
}
