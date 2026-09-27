package repository

import "employee-management/model"

type EmployeeRepository interface {
	AddEmployee(employee model.Employee) error
	GetAllEmployees() []model.Employee
	GetEmployeeByID(id int) model.Employee
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
}