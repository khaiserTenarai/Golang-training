package service

import (
	"context"
	"fmt"
	"strings"

	"department-management/model"
	"department-management/repository"
	"department-management/utility"
)

type EmployeeService struct {
	EmployeeRepository   *repository.EmployeeRepository
	DepartmentRepository *repository.DepartmentRepository
}

func NewEmployeeService(
	employeeRepository *repository.EmployeeRepository,
	departmentRepository *repository.DepartmentRepository,
) *EmployeeService {

	return &EmployeeService{
		EmployeeRepository:   employeeRepository,
		DepartmentRepository: departmentRepository,
	}
}

func (s *EmployeeService) Create(
	ctx context.Context,
	name string,
	email string,
	salary float64,
	departmentID int,
) (*model.Employee, error) {

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if !utility.ValidateEmployeeName(name) {
		return nil, fmt.Errorf(
			"employee name must contain at least 2 characters",
		)
	}

	if !utility.ValidateEmail(email) {
		return nil, fmt.Errorf(
			"invalid email address",
		)
	}

	if !utility.ValidateSalary(salary) {
		return nil, fmt.Errorf(
			"salary cannot be negative",
		)
	}

	if !utility.ValidateID(departmentID) {
		return nil, fmt.Errorf(
			"invalid department ID",
		)
	}

	_, err := s.DepartmentRepository.GetByID(
		ctx,
		departmentID,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"department does not exist",
		)
	}

	employee := &model.Employee{
		Name:         name,
		Email:        email,
		Salary:       salary,
		DepartmentID: departmentID,
	}

	err = s.EmployeeRepository.Create(
		ctx,
		employee,
	)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *EmployeeService) GetByID(
	ctx context.Context,
	id int,
) (*model.EmployeeWithDepartment, error) {

	if !utility.ValidateID(id) {
		return nil, fmt.Errorf("invalid employee ID")
	}

	return s.EmployeeRepository.GetByID(
		ctx,
		id,
	)
}

func (s *EmployeeService) GetAll(
	ctx context.Context,
) ([]model.EmployeeWithDepartment, error) {

	return s.EmployeeRepository.GetAll(ctx)
}

func (s *EmployeeService) Update(
	ctx context.Context,
	id int,
	name string,
	email string,
	salary float64,
	departmentID int,
) error {

	if !utility.ValidateID(id) {
		return fmt.Errorf("invalid employee ID")
	}

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if !utility.ValidateEmployeeName(name) {
		return fmt.Errorf(
			"employee name must contain at least 2 characters",
		)
	}

	if !utility.ValidateEmail(email) {
		return fmt.Errorf(
			"invalid email address",
		)
	}

	if !utility.ValidateSalary(salary) {
		return fmt.Errorf(
			"salary cannot be negative",
		)
	}

	if !utility.ValidateID(departmentID) {
		return fmt.Errorf(
			"invalid department ID",
		)
	}

	_, err := s.DepartmentRepository.GetByID(
		ctx,
		departmentID,
	)

	if err != nil {
		return fmt.Errorf(
			"department does not exist",
		)
	}

	employee := &model.Employee{
		ID:           id,
		Name:         name,
		Email:        email,
		Salary:       salary,
		DepartmentID: departmentID,
	}

	return s.EmployeeRepository.Update(
		ctx,
		employee,
	)
}

func (s *EmployeeService) Delete(
	ctx context.Context,
	id int,
) error {

	if !utility.ValidateID(id) {
		return fmt.Errorf("invalid employee ID")
	}

	return s.EmployeeRepository.Delete(
		ctx,
		id,
	)
}

func (s *EmployeeService) ChangeDepartment(
	ctx context.Context,
	employeeID int,
	departmentID int,
) error {

	if !utility.ValidateID(employeeID) {
		return fmt.Errorf("invalid employee ID")
	}

	if !utility.ValidateID(departmentID) {
		return fmt.Errorf("invalid department ID")
	}

	_, err := s.EmployeeRepository.GetByID(
		ctx,
		employeeID,
	)

	if err != nil {
		return fmt.Errorf(
			"employee does not exist",
		)
	}

	_, err = s.DepartmentRepository.GetByID(
		ctx,
		departmentID,
	)

	if err != nil {
		return fmt.Errorf(
			"department does not exist",
		)
	}

	return s.EmployeeRepository.ChangeDepartment(
		ctx,
		employeeID,
		departmentID,
	)
}
