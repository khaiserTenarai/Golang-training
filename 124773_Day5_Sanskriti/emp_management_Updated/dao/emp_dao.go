package dao

import (
	"errors"

	"emp_management_Updated/model"
)

var ErrEmployeeNotFound = errors.New("employee not found")

// DAO interface
type EmployeeDAO interface {
	AddEmployee(employee *model.Employee)
	GetEmployees() []model.Employee
	DeleteEmployee(id int) bool
	UpdateEmployee(employee *model.Employee) bool
}

// DAO implementation
type EmployeeDAOImpl struct {
	employees []model.Employee
}

// Constructor
func NewEmployeeDAO() EmployeeDAO {
	return &EmployeeDAOImpl{}
}

// Add employee
func (d *EmployeeDAOImpl) AddEmployee(employee *model.Employee) {
	d.employees = append(d.employees, *employee)
}

// Get employees
func (d *EmployeeDAOImpl) GetEmployees() []model.Employee {
	return d.employees
}

// Delete employee
func (d *EmployeeDAOImpl) DeleteEmployee(id int) bool {

	for i, employee := range d.employees {

		if employee.ID == id {

			d.employees = append(
				d.employees[:i],
				d.employees[i+1:]...,
			)

			return true
		}
	}

	return false
}

// Update employee
func (d *EmployeeDAOImpl) UpdateEmployee(employee *model.Employee) bool {

	for i := range d.employees {

		if d.employees[i].ID == employee.ID {

			d.employees[i] = *employee

			return true
		}
	}

	return false
}
