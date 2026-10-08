package controller

import "employee-management/model"

type DepartmentController interface {
	CreateDepartment(department model.Department) error
	GetDepartment(id int) (*model.Department, error)
	GetAllDepartments() ([]model.Department, error)
	UpdateDepartment(department model.Department) error
	DeleteDepartment(id int) error
}
