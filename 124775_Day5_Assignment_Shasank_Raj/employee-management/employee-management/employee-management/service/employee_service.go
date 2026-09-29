package service

import "employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee) error
	GetEmployee(id int) (model.Employee, error)
	GetAllEmployees() []model.Employee
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
	SearchEmployee(name string) []model.Employee
}

