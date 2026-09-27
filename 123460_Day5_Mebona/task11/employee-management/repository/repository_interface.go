package repository

import "employee-management/model"

type EmployeeRepository interface {
	Add(employee model.Employee)
	Get(id int) (model.Employee, bool)
	GetAll() []model.Employee
	Update(employee model.Employee) bool
	Delete(id int) bool
}