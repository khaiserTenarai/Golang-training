package repository

import "example.com/employee-management/model"

type DepartmentRepository interface {
	Save(dept *model.Department) error
	FindByID(id int64) (model.Department, error)
	FindAll() ([]model.Department, error)
	Update(dept model.Department) error
	Delete(id int64) error
}