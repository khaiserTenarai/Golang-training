package service

import "employee-management-app/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee) error

	DeleteEmployee(id int) error

	UpdateEmployee(employee model.Employee) error

	FindEmployeeByID(id int) (model.Employee, error)

	FindAllEmployees() []model.Employee

	Search(name string, department string, minSalary float64) []model.Employee

	UpdateSalary(employeeID int, newSalary float64) error
}
