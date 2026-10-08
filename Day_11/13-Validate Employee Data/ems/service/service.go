package service

import (
	"errors"

	"employee-management-go/model"
	"employee-management-go/repository"
	"employee-management-go/utility"
)

// EmployeeService defines business operations.
type EmployeeService interface {
	Create(employee model.Employee) (model.Employee, error)
	GetAll() []model.Employee
	GetByID(id int) (model.Employee, error)
	Delete(id int) error
}

// EmployeeServiceImpl implements EmployeeService.
type EmployeeServiceImpl struct {
	Repository repository.EmployeeRepository
}

// Create validates and creates an employee.
func (s *EmployeeServiceImpl) Create(employee model.Employee) (model.Employee, error) {

	// Validate employee data.
	err := utility.ValidateEmployee(employee)

	if err != nil {
		return employee, err
	}

	return s.Repository.Create(employee)
}

// GetAll returns all employees.
func (s *EmployeeServiceImpl) GetAll() []model.Employee {
	return s.Repository.GetAll()
}

// GetByID returns one employee.
func (s *EmployeeServiceImpl) GetByID(id int) (model.Employee, error) {

	if id <= 0 {
		return model.Employee{}, errors.New("invalid employee ID")
	}

	return s.Repository.GetByID(id)
}

// Delete deletes an employee.
func (s *EmployeeServiceImpl) Delete(id int) error {

	if id <= 0 {
		return errors.New("invalid employee ID")
	}

	return s.Repository.Delete(id)
}
