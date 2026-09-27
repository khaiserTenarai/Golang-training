package repository

import "employee-management/model"

type EmployeeRepository interface {
	AddEmployee(employee model.Employee) error
	GetEmployees() []model.Employee
	GetEmployee(id int) (model.Employee, error)
	DeleteEmployee(id int) error
}