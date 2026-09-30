package service

import "department-management/model"

type EmployeeService interface {
	Save(employee model.Employee) error
	FindByID(id int) (model.Employee, error)
	FindAll() ([]model.Employee, error)
	Update(employee model.Employee) error
	Delete(id int) error
}
