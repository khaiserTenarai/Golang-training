package service

import "day7q12/model"

type EmployeeService interface {
	CreateEmployee(employee model.Employee) error

	GetEmployees() ([]model.Employee, error)
}
