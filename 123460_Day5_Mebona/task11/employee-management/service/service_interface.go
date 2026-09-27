package service

import "employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee) bool
	GetEmployee(id int) (model.Employee, bool)
	GetAllEmployees() []model.Employee
	UpdateEmployee(employee model.Employee) bool
	DeleteEmployee(id int) bool
}