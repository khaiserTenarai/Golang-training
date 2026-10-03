package repository

import "assignment_task12/model"

type EmployeeRepository interface {
	Create(employee model.Employee) error

	GetAll() ([]model.Employee, error)
}
