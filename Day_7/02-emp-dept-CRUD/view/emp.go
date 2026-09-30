package view

import "department-management/model"

type EmployeeView interface {
	ShowMenu() int
	ReadEmployee() model.Employee
	ReadEmployeeForUpdate() model.Employee
	ReadID() int
	DisplayEmployee(employee model.Employee)
	DisplayEmployees(employees []model.Employee)
}
