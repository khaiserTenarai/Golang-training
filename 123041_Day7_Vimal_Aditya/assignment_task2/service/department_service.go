package service

import "example.com/employee-management/model"

type DepartmentService interface {
	AddDepartment(department *model.Department) error
	GetDepartment(id int64) (model.Department, error)
	GetAllDepartments() ([]model.Department, error)
	UpdateDepartment(department model.Department) error
	DeleteDepartment(id int64) error
}
