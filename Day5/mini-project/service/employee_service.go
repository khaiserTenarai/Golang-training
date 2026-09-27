package service

import "example.com/employee-management/model"


type EmployeeService interface {
	AddEmployee(employee model.Employee) (model.Employee, error)
	GetEmployee(id int) (model.Employee, error)
	GetAllEmployees() []model.Employee
	UpdateEmployee(employee model.Employee) (model.Employee, error)
	DeleteEmployee(id int) error
}
