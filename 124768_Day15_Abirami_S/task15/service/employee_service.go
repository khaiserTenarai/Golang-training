package service

import "day15-secure-gin-api/model"

type EmployeeService interface {
	GetAll(int, int) []model.Employee
	GetByID(int) (model.Employee, bool)
	Create(model.Employee) model.Employee
	Update(int, model.Employee) (model.Employee, bool)
	Delete(int) bool
}
