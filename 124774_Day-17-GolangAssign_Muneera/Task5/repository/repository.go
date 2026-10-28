package repository

import (
	"errors"

	"employee-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(*model.Employee) error
	GetByID(uint) (*model.Employee, error)
	GetAll() ([]model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(e *model.Employee) error {
	// Check that the referenced department exists.
	var department model.Department
	if err := r.db.First(&department, e.DepartmentID).Error; err != nil {
		return err
	}

	return r.db.Create(e).Error
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee

	err := r.db.
		Preload("Department").
		First(&employee, id).Error

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	var employees []model.Employee

	err := r.db.
		Preload("Department").
		Order("id ASC").
		Find(&employees).Error

	return employees, err
}

func (r *employeeRepository) Update(
	id uint,
	updated *model.Employee,
) (*model.Employee, error) {
	var employee model.Employee

	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	var department model.Department
	if err := r.db.First(&department, updated.DepartmentID).Error; err != nil {
		return nil, err
	}

	employee.Name = updated.Name
	employee.Email = updated.Email
	employee.DepartmentID = updated.DepartmentID
	employee.Salary = updated.Salary

	if err := r.db.Save(&employee).Error; err != nil {
		return nil, err
	}

	// Load the related department for the returned result.
	if err := r.db.Preload("Department").
		First(&employee, employee.ID).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Employee{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}

	return nil
}
