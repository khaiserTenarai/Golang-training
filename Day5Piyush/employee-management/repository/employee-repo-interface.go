package repository

import "Day5Piyush/employee-management/model"

type EmployeeRepository interface {
	Add(employee model.Employee) error
	GetByID(id int) (model.Employee, error)
	GetAll() []model.Employee
	Delete(id int) error
}
