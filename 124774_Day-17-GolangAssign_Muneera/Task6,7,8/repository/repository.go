package repository

import (
	"employee-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(e *model.Employee) error
	GetByID(id uint) (*model.Employee, error)
	GetAll() ([]model.Employee, error)
	Update(id uint, e *model.Employee) (*model.Employee, error)
	Delete(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(e *model.Employee) error {
	return r.db.Create(e).Error
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	var employees []model.Employee

	err := r.db.Order("id").Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee

	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) Update(
	id uint,
	e *model.Employee,
) (*model.Employee, error) {
	var employee model.Employee

	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	employee.Name = e.Name
	employee.Email = e.Email
	employee.Department = e.Department
	employee.Salary = e.Salary

	if err := r.db.Save(&employee).Error; err != nil {
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
		return gorm.ErrRecordNotFound
	}

	return nil
}
