package service

import "example.com/employee-management/models"

type EmployeeService interface {
	AddEmployee(employee models.Employee) error
	GetEmployee(id int64) (models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
	UpdateEmployee(employee models.Employee) error
	DeleteEmployee(id int64) error

	// Task 2: Department Mapping
	AssignDepartment(employeeID int64, departmentID int) error
	GetEmployeeWithDepartment(id int64) (models.EmployeeWithDepartment, error)
}