package view

import "employee-management/model"

type EmployeeView interface {
	Start()
	CreateEmployee()
	GetEmployee()
	GetAllEmployees()
	UpdateEmployee()
	DeleteEmployee()
	SearchEmployees()
	DisplayEmployee(employee *model.Employee)
	DisplayEmployees(employees []model.Employee)
}
