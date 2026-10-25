package service

import (
	"errors"
	"net/mail"
	"strings"

	"ems/model"
	"ems/repository"

	"gorm.io/gorm"
)

type EmployeeService interface {
	Create(*model.Employee) error
	GetAll(page, pageSize int) ([]model.Employee, int64, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: r}
}

// Validate employee data before database operations.
func validate(e *model.Employee) error {
	if e == nil {
		return errors.New("employee is required")
	}

	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)

	if e.Name == "" || len(e.Name) > 100 {
		return errors.New("name is required and must be at most 100 characters")
	}

	address, err := mail.ParseAddress(e.Email)
	if err != nil || address.Address != e.Email || len(e.Email) > 150 {
		return errors.New("a valid email is required")
	}

	if e.Department == "" || len(e.Department) > 100 {
		return errors.New("department is required and must be at most 100 characters")
	}

	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}

	return nil
}

func (s *employeeService) Create(e *model.Employee) error {
	if err := validate(e); err != nil {
		return err
	}

	// Let PostgreSQL generate the ID.
	e.ID = 0
	return s.repo.Create(e)
}

func (s *employeeService) GetAll(
	page, pageSize int,
) ([]model.Employee, int64, error) {
	if page < 1 {
		return nil, 0, errors.New("page must be at least 1")
	}

	if pageSize < 1 || pageSize > 100 {
		return nil, 0, errors.New("pageSize must be between 1 and 100")
	}

	return s.repo.GetAll(page, pageSize)
}

func (s *employeeService) GetByID(id uint) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("invalid employee ID")
	}
	return s.repo.GetByID(id)
}

func (s *employeeService) Update(
	id uint,
	e *model.Employee,
) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("invalid employee ID")
	}

	if err := validate(e); err != nil {
		return nil, err
	}

	return s.repo.Update(id, e)
}

func (s *employeeService) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid employee ID")
	}
	return s.repo.Delete(id)
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
