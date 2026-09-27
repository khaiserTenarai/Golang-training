package repository

import "example.com/employee-management/model"

type EmployeeRepository interface {
	Save(employee model.Employee) model.Employee
	FindByID(id int) (model.Employee, bool)
	FindAll() []model.Employee
	Update(employee model.Employee) (model.Employee, bool)
	Delete(id int) bool
}
