package repository

import "employee-management/model"

type EmployeeRepository interface {
	Add(employee model.Employee)
	GetByID(id int) (model.Employee, bool)
	GetAll() []model.Employee
	Delete(id int)
}
