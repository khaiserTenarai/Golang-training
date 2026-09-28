package controller

import "employee-management/model"

type EmployeeController interface {
	Add(employee model.Employee) error
	GetEmployeeById(id int) (model.Employee, error)
	GetEmployee() []model.Employee
	UpdateEmp(employee model.Employee) error
	DeleteEmp(id int) error
}
