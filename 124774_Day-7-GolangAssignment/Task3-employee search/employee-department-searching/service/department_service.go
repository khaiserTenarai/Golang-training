package service

import "employee-management-app/model"

type DepartmentService interface {
	AddDepartment(department model.Department) error

	DeleteDepartment(id int) error

	UpdateDepartment(department model.Department) error

	FindDepartmentByID(id int) (model.Department, error)

	FindAllDepartments() []model.Department
}
