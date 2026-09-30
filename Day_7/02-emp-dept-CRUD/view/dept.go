package view

import "department-management/model"

type DepartmentView interface {
	ShowMenu() int
	ReadDepartment() model.Department
	ReadDepartmentForUpdate() model.Department
	ReadID() int
	DisplayDepartment(department model.Department)
	DisplayDepartments(departments []model.Department)
}
