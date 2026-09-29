package service

import "employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee)
	GetEmployee(id int) (model.Employee, bool)
	GetAllEmployees() []model.Employee
	DeleteEmployee(id int)
}
