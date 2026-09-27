package service

import "example.com/employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee)
	GetAllEmployees() []model.Employee
}
