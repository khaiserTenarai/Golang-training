package service

import (
	"errors"
	"strings"

	"employee-platform/model"
	"employee-platform/repository"
	"gorm.io/gorm"
)

type EmployeeService struct {
	employees   *repository.EmployeeRepository
	departments *repository.DepartmentRepository
}

func NewEmployeeService(e *repository.EmployeeRepository, d *repository.DepartmentRepository) *EmployeeService {
	return &EmployeeService{employees: e, departments: d}
}

type EmployeeList struct {
	Items []model.Employee `json:"items"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}

func (s *EmployeeService) List(page, limit int, departmentID uint, search, sortBy, order string) (*EmployeeList, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	items, total, err := s.employees.List(page, limit, departmentID, strings.TrimSpace(search), sortBy, strings.ToLower(order))
	if err != nil {
		return nil, err
	}
	return &EmployeeList{Items: items, Total: total, Page: page, Limit: limit}, nil
}
func (s *EmployeeService) ByID(id uint) (*model.Employee, error) { return s.employees.ByID(id) }
func (s *EmployeeService) Create(name, email string, salary float64, departmentID uint) (*model.Employee, error) {
	name, email = strings.TrimSpace(name), strings.ToLower(strings.TrimSpace(email))
	if name == "" || email == "" || salary < 0 || departmentID == 0 {
		return nil, errors.New("name, email, non-negative salary and department_id are required")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("email is invalid")
	}
	if !s.departments.Exists(departmentID) {
		return nil, errors.New("department does not exist")
	}
	e := &model.Employee{Name: name, Email: email, Salary: salary, DepartmentID: departmentID}
	if err := s.employees.Create(e); err != nil {
		return nil, err
	}
	return s.employees.ByID(e.ID)
}
func (s *EmployeeService) Update(id uint, name, email string, salary float64, departmentID uint) (*model.Employee, error) {
	e, err := s.employees.ByID(id)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(name) != "" {
		e.Name = strings.TrimSpace(name)
	}
	if strings.TrimSpace(email) != "" {
		e.Email = strings.ToLower(strings.TrimSpace(email))
	}
	if salary < 0 {
		return nil, errors.New("salary cannot be negative")
	}
	e.Salary = salary
	if departmentID > 0 {
		if !s.departments.Exists(departmentID) {
			return nil, errors.New("department does not exist")
		}
		e.DepartmentID = departmentID
	}
	if err := s.employees.Update(e); err != nil {
		return nil, err
	}
	return s.employees.ByID(id)
}
func (s *EmployeeService) Delete(id uint) error {
	if _, err := s.employees.ByID(id); err != nil {
		return err
	}
	return s.employees.Delete(id)
}
func IsNotFound(err error) bool { return errors.Is(err, gorm.ErrRecordNotFound) }
