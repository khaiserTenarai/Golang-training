package service

import "example.com/employee-management/models"

type DepartmentService interface {
	AddDepartment(dept models.Department) error
	GetDepartment(id int) (models.Department, error)
	GetAllDepartments() ([]models.Department, error)
	UpdateDepartment(dept models.Department) error
	DeleteDepartment(id int) error
}