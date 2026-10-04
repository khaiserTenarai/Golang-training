package repository

import "employee-management-app/model"

type DepartmentRepository interface {
	Save(department model.Department) error

	FindByID(id int) (model.Department, error)

	FindAll() []model.Department

	Update(department model.Department) error

	Delete(id int) error
}
