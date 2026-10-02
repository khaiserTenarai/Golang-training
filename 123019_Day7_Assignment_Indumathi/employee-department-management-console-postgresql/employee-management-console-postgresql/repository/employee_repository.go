package repository

import (
	"example.com/employee-management/models"
)

type EmployeeRepository interface {
	Save(employee models.Employee) error
	FindByID(id int64) (models.Employee, error)
	FindAll() ([]models.Employee, error)
	Update(employee models.Employee) error
	Delete(id int64) error
	
	// Task 2: Department Mapping methods
	AssignDepartment(employeeID int64, departmentID int) error
	FindWithDepartment(id int64) (models.EmployeeWithDepartment, error)
}