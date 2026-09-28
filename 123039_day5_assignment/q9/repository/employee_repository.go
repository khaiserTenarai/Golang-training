package repository

import "example.com/employee_management/model"

type EmployeeRepository interface {
	AddEmployee(employee model.Employee)
	GetEmployee(id int) (model.Employee, error)
	GetAllEmployees() []model.Employee
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
}
