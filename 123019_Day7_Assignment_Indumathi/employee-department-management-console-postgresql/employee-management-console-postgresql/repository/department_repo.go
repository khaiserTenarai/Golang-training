package repository

import "example.com/employee-management/models"

type DepartmentRepository interface {
	Create(dept *models.Department) error
	FindByID(id int) (*models.Department, error)
	FindAll() ([]models.Department, error)
	Update(dept *models.Department) error
	Delete(id int) error
}