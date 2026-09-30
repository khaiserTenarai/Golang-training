package service

import "department-management/model"

type DepartmentService interface {
	Save(department model.Department) error
	FindByID(id int) (model.Department, error)
	FindAll() ([]model.Department, error)
	Update(department model.Department) error
	Delete(id int) error
}
