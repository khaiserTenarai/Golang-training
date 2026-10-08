package repository

import "employee-management/model"

type EmployeeRepository interface {
	CreateEmployee(employee model.Employee) error
	GetEmployee(id int) (*model.Employee, error)
	GetAllEmployees() ([]model.Employee, error)
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
	SearchEmployees(name string, departmentID int, salary float64, page int, size int, sortBy string, sortOrder string) ([]model.Employee, error)
}
