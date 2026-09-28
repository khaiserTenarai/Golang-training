package service

import "employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee) error
	GetById(ID int) (model.Employee, error)
	GetAll() []model.Employee
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(ID int) error
}
