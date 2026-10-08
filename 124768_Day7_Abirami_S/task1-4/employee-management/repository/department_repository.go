package repository

import "employee-management/model"

type DepartmentRepository interface {
	CreateDepartment(department model.Department) error
	GetDepartment(id int) (*model.Department, error)
	GetAllDepartments() ([]model.Department, error)
	UpdateDepartment(department model.Department) error
	DeleteDepartment(id int) error
}
