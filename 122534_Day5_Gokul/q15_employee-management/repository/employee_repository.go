package repository

import "example.com/employee-management/model"

type EmployeeRepository interface {
	Save(employee model.Employee)
	FindAll() []model.Employee
}
