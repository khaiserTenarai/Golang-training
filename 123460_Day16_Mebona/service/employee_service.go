package service

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"employee-management/apperror"
	"employee-management/model"
	"employee-management/repository"
)

const maxPageSize = 100

type EmployeeService interface {
	Get(id uint) (*model.Employee, error)
	List(f model.EmployeeFilter) ([]*model.Employee, error)
	Create(in model.CreateEmployeeInput) (*model.Employee, error)
	Update(id uint, in model.UpdateEmployeeInput) (*model.Employee, error)
	Delete(id uint) error
}

type employeeService struct {
	employees repository.EmployeeRepository
	depts     repository.DepartmentRepository
}

func NewEmployeeService(e repository.EmployeeRepository, d repository.DepartmentRepository) EmployeeService {
	return &employeeService{employees: e, depts: d}
}

func isNotFound(err error) bool { return errors.Is(err, repository.ErrNotFound) }

func (s *employeeService) Get(id uint) (*model.Employee, error) {
	e, err := s.employees.GetByID(id)
	if isNotFound(err) {
		return nil, apperror.NotFound(fmt.Sprintf("employee with id %d not found", id))
	}
	return e, err
}

func (s *employeeService) List(f model.EmployeeFilter) ([]*model.Employee, error) {
	if f.Limit < 1 || f.Limit > maxPageSize {
		return nil, apperror.BadInput("limit", fmt.Sprintf("limit must be between 1 and %d", maxPageSize))
	}
	if f.Offset < 0 {
		return nil, apperror.BadInput("offset", "offset must be 0 or greater")
	}
	return s.employees.List(f)
}

func (s *employeeService) Create(in model.CreateEmployeeInput) (*model.Employee, error) {
	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.TrimSpace(in.Email)
	if err := s.validate(in.Name, in.Email, in.Salary, in.DepartmentID); err != nil {
		return nil, err
	}
	if err := s.ensureEmailFree(in.Email, 0); err != nil {
		return nil, err
	}
	active := true
	if in.Active != nil {
		active = *in.Active
	}
	e := &model.Employee{
		Name: in.Name, Email: in.Email, Salary: in.Salary, Active: active,
		DepartmentID: in.DepartmentID, JoinedAt: time.Now().UTC(),
	}
	if err := s.employees.Create(e); err != nil {
		return nil, err
	}
	return s.employees.GetByID(e.ID) // reload so the department is filled in
}

// Update applies only the fields the client actually sent.
func (s *employeeService) Update(id uint, in model.UpdateEmployeeInput) (*model.Employee, error) {
	e, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	if in.Name != nil {
		e.Name = strings.TrimSpace(*in.Name)
	}
	if in.Email != nil {
		e.Email = strings.TrimSpace(*in.Email)
	}
	if in.Salary != nil {
		e.Salary = *in.Salary
	}
	if in.DepartmentID != nil {
		e.DepartmentID = *in.DepartmentID
	}
	if in.Active != nil {
		e.Active = *in.Active
	}
	if err := s.validate(e.Name, e.Email, e.Salary, e.DepartmentID); err != nil {
		return nil, err
	}
	if err := s.ensureEmailFree(e.Email, e.ID); err != nil {
		return nil, err
	}
	if err := s.employees.Update(e); err != nil {
		return nil, err
	}
	return s.employees.GetByID(e.ID)
}

func (s *employeeService) Delete(id uint) error {
	err := s.employees.Delete(id)
	if isNotFound(err) {
		return apperror.NotFound(fmt.Sprintf("employee with id %d not found", id))
	}
	return err
}

func (s *employeeService) validate(name, email string, salary float64, deptID uint) error {
	if len(name) < 2 || len(name) > 50 {
		return apperror.BadInput("name", "name must be between 2 and 50 characters")
	}
	if addr, err := mail.ParseAddress(email); err != nil || addr.Address != email {
		return apperror.BadInput("email", "email is not a valid address")
	}
	if salary <= 0 {
		return apperror.BadInput("salary", "salary must be greater than 0")
	}
	if _, err := s.depts.GetByID(deptID); err != nil {
		if isNotFound(err) {
			return apperror.BadInput("departmentId", fmt.Sprintf("department %d does not exist", deptID))
		}
		return err
	}
	return nil
}

// ensureEmailFree fails with CONFLICT if another employee (not selfID) uses the email.
func (s *employeeService) ensureEmailFree(email string, selfID uint) error {
	other, err := s.employees.GetByEmail(email)
	if isNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if other.ID != selfID {
		return apperror.Conflict("email", fmt.Sprintf("email %q is already in use", email))
	}
	return nil
}
