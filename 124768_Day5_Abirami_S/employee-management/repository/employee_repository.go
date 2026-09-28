package repository

import "employee-management/model"

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindById(ID int) (model.Employee, error)
	FindAll() []model.Employee
	Update(employee model.Employee) error
	Delete(ID int) error
}
