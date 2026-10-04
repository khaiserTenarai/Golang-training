package repository

import "employee-app/model"

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindAll() []model.Employee
}
