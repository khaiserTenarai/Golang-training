package controller

import (
	"emp_management_Updated/model"
	"emp_management_Updated/service"
)

// Controller interface
type EmployeeController interface {
	AddEmployee(employee *model.Employee) error
	DisplayEmployees() []model.Employee
	DeleteEmployee(id int) bool
	UpdateEmployee(employee *model.Employee) error
}

// Controller implementation
type EmployeeControllerImpl struct {
	service service.EmployeeService
}

// Constructor
func NewEmployeeController(s service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		service: s,
	}
}

// Add employee
func (c *EmployeeControllerImpl) AddEmployee(employee *model.Employee) error {
	return c.service.AddEmployee(employee)
}

// Display employees
func (c *EmployeeControllerImpl) DisplayEmployees() []model.Employee {
	return c.service.GetEmployees()
}

// Delete employee
func (c *EmployeeControllerImpl) DeleteEmployee(id int) bool {
	return c.service.DeleteEmployee(id)
}

// Update employee
func (c *EmployeeControllerImpl) UpdateEmployee(employee *model.Employee) error {
	return c.service.UpdateEmployee(employee)
}
