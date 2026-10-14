package service

import "employeegin/model"

type EmployeeService interface {
	AddEmployee(req model.EmployeeRequest) (model.Employee, error)
	GetEmployee(id int) (model.Employee, error)
	GetAllEmployees(department string) ([]model.Employee, error)
	ReplaceEmployee(id int, req model.EmployeeRequest) (model.Employee, error)
	PatchEmployee(id int, patch model.EmployeePatch) (model.Employee, error)
	DeleteEmployee(id int) error
}
