package service

import (
	"emp_management_Updated/dao"
	"emp_management_Updated/model"
	"emp_management_Updated/utility"
)

// Service interface
type EmployeeService interface {
	AddEmployee(employee *model.Employee) error
	GetEmployees() []model.Employee
	DeleteEmployee(id int) bool
	UpdateEmployee(employee *model.Employee) error
}

// Service implementation
type EmployeeServiceImpl struct {
	dao dao.EmployeeDAO
}

// Constructor
func NewEmployeeService(d dao.EmployeeDAO) EmployeeService {
	return &EmployeeServiceImpl{
		dao: d,
	}
}

// Add employee
func (s *EmployeeServiceImpl) AddEmployee(employee *model.Employee) error {

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateEntry(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	s.dao.AddEmployee(employee)

	return nil
}

// Get employees
func (s *EmployeeServiceImpl) GetEmployees() []model.Employee {
	return s.dao.GetEmployees()
}

// Delete employee
func (s *EmployeeServiceImpl) DeleteEmployee(id int) bool {
	return s.dao.DeleteEmployee(id)
}

// Update employee
func (s *EmployeeServiceImpl) UpdateEmployee(employee *model.Employee) error {

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateEntry(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	if !s.dao.UpdateEmployee(employee) {
		return dao.ErrEmployeeNotFound
	}

	return nil
}
