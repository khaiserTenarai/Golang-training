package service

import "employee_management/models"

type EmployeeService interface{
	AddEmployee(employee models.Employee) error
	GetEmployee(id int) (models.Employee, error)
	GetAllEmployee() ([]models.Employee, error)
	UpdateEmployee(employee models.Employee) error
	DeleteEmployee(id int) error
}