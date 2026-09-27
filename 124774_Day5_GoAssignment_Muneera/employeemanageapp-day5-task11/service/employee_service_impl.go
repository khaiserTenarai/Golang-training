package service

import (
	"errors"

	"employee-management-app/model"
	"employee-management-app/repository"
	"employee-management-app/utility"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

// Creates Service and receives Repository.
func NewEmployeeService(
	repository repository.EmployeeRepository,
) EmployeeService {

	return &EmployeeServiceImpl{
		repository: repository,
	}
}

// Add employee.
func (s *EmployeeServiceImpl) AddEmployee(
	employee model.Employee,
) error {

	// Validation happens in Service.

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateEmail(employee.Email); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	// After validation, send to Repository.
	return s.repository.Save(employee)
}

// Delete employee.
func (s *EmployeeServiceImpl) DeleteEmployee(
	id int,
) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	return s.repository.Delete(id)
}

// Update employee.
func (s *EmployeeServiceImpl) UpdateEmployee(
	employee model.Employee,
) error {

	// Validate updated employee.

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateEmail(employee.Email); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	return s.repository.Update(employee)
}

// Find employee by ID.
func (s *EmployeeServiceImpl) FindEmployeeByID(
	id int,
) (model.Employee, error) {

	if id <= 0 {
		return model.Employee{}, errors.New(
			"ID must be greater than 0",
		)
	}

	return s.repository.FindByID(id)
}

// Find all employees.
func (s *EmployeeServiceImpl) FindAllEmployees() []model.Employee {

	return s.repository.FindAll()
}
