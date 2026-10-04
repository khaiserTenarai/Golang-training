package service

import "employee-app/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee) error
	GetEmployees() []model.Employee
}
