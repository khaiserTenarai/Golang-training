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

func (r *employeeRepository) Create(employee *model.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	var employees []model.Employee
	err := r.db.Order("id ASC").Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) Update(id uint, updated *model.Employee) (*model.Employee, error) {
	employee, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	employee.Name = updated.Name
	employee.Email = updated.Email
	employee.Department = updated.Department
	employee.Salary = updated.Salary

	if err := r.db.Save(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
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
