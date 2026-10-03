package view

import "employee-management/model"

type DepartmentView interface {
	Start()
	CreateDepartment()
	GetDepartment()
	GetAllDepartments()
	UpdateDepartment()
	DeleteDepartment()
	DisplayDepartment(department model.Department)
	DisplayDepartments(departments []model.Department)
}
