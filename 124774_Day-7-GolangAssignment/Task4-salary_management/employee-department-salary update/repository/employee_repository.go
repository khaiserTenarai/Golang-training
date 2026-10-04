package repository

import "employee-management-app/model"

type EmployeeRepository interface {
	Save(employee model.Employee) error

	Delete(id int) error

	Update(employee model.Employee) error

	FindByID(id int) (model.Employee, error)

	FindAll() []model.Employee

	Search(name string, department string, minSalary float64) []model.Employee

	UpdateSalary(employeeID int, newSalary float64) error
}
