package service

import "assignment_task12/model"

type EmployeeService interface {
	CreateEmployee(employee model.Employee) error

	GetEmployees() ([]model.Employee, error)
}
