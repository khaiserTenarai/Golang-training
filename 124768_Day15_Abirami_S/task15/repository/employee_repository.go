package repository

import "day15-secure-gin-api/model"

type EmployeeRepository interface {
	GetAll() []model.Employee
	GetByID(int) (model.Employee, bool)
	Create(model.Employee) model.Employee
	Update(int, model.Employee) (model.Employee, bool)
	Delete(int) bool
}
