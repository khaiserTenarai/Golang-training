package repository

import "day7q12/model"

type EmployeeRepository interface {
	Create(employee model.Employee) error

	GetAll() ([]model.Employee, error)
}
