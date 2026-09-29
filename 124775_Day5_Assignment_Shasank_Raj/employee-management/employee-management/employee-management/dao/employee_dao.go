package dao

import "employee-management/model"

type EmployeeDAO interface {
	AddEmployee(employee model.Employee) error
	GetEmployeeByID(id int) (model.Employee, error)
	GetAllEmployees() []model.Employee
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
	SearchEmployee(name string) []model.Employee
}
