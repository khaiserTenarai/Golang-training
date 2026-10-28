package repository

import "day16-employee-graphql/model"

type EmployeeRepository interface {
	GetByID(id int) (*model.Employee, error)
	GetAll() ([]model.Employee, error)
	Create(employee model.Employee) (*model.Employee, error)
	Update(id int, employee model.Employee) (*model.Employee, error)
	Delete(id int) (*model.Employee, error)
}
