package view
import "employee-management/model"

type DepartmentView interface {
	ShowMenu() int
	ReadDepartment() model.Department
	ReadID() int
	DisplayDepartment(department model.Department)
	DisplayDepartments(departments []model.Department)
}
