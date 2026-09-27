package dao

import "employee_management/models"

type EmployeeDao interface{
	Save(employee models.Employee) error
	FindById(id int) (models.Employee, error)
	FindAll() ([]models.Employee, error)
	Update(employee models.Employee) error
	Delete(id int) error
}