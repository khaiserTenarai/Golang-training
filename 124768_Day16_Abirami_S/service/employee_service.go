package service

import "day16-employee-graphql/model"

type EmployeeService interface {
	GetEmployee(id int) (*model.Employee, error)
	GetEmployees() ([]model.Employee, error)
	CreateEmployee(employee model.Employee) (*model.Employee, error)
	UpdateEmployee(id int, employee model.Employee) (*model.Employee, error)
	DeleteEmployee(id int) (*model.Employee, error)
}
