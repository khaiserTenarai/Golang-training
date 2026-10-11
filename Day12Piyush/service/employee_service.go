// Tasks 2, 3: Service Layer — business logic between handler and repository.
// Task 4: Depends on repository INTERFACE (dependency injection).

package service

import (
	"production_api/models"
	"production_api/repository"
	"production_api/validator"
)

// ---- Service Interface (Task 4: DI contract for handler) ----

type EmployeeService interface {
	Create(emp models.Employee) (models.Employee, []models.ValidationError, error)
	GetAll(params models.QueryParams) ([]models.Employee, int, error)
	GetByID(id int) (models.Employee, error)
	Update(id int, emp models.Employee) (models.Employee, []models.ValidationError, error)
	Delete(id int) error
	Count() int
}

// ---- Implementation ----

type employeeServiceImpl struct {
	repo repository.EmployeeRepository // depends on interface, not concrete type
}

// NewEmployeeService injects the repository dependency (Task 4)
func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo: repo}
}

func (s *employeeServiceImpl) Create(emp models.Employee) (models.Employee, []models.ValidationError, error) {
	// Task 9: Validation layer
	if errs := validator.ValidateEmployee(emp); len(errs) > 0 {
		return models.Employee{}, errs, nil
	}
	created, err := s.repo.Create(emp)
	return created, nil, err
}

func (s *employeeServiceImpl) GetAll(params models.QueryParams) ([]models.Employee, int, error) {
	return s.repo.GetAll(params)
}

func (s *employeeServiceImpl) GetByID(id int) (models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeServiceImpl) Update(id int, emp models.Employee) (models.Employee, []models.ValidationError, error) {
	if errs := validator.ValidateEmployee(emp); len(errs) > 0 {
		return models.Employee{}, errs, nil
	}
	updated, err := s.repo.Update(id, emp)
	return updated, nil, err
}

func (s *employeeServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *employeeServiceImpl) Count() int {
	return s.repo.Count()
}
