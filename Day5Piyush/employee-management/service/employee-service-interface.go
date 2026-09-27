package service

import "Day5Piyush/employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee) error
	GetEmployee(id int) (model.Employee, error)
	GetEmployees() []model.Employee
	DeleteEmployee(id int) error
}
