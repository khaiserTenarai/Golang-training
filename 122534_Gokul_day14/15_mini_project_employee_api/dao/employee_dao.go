package dao

import "employeegin/model"

type EmployeeDAO interface {
	Save(employee model.Employee) (model.Employee, error)
	FindByID(id int) (model.Employee, error)
	FindByEmail(email string) (model.Employee, error)
	FindAll(department string) ([]model.Employee, error)
	Update(employee model.Employee) error
	Delete(id int) error
}
