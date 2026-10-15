package repository

import (
	"employee-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(id uint) (*model.Employee, error)
	Update(id uint, employee model.Employee) (*model.Employee, error)
	Delete(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

// NewEmployeeRepository creates a repository using the database connection.
func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(employee *model.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	var employees []model.Employee

	err := r.db.Order("id ASC").Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee

	err := r.db.First(&employee, id).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) Update(
	id uint,
	updated model.Employee,
) (*model.Employee, error) {
	var employee model.Employee

	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	// A map allows zero values, such as salary = 0, to be updated.
	err := r.db.Model(&employee).Updates(map[string]interface{}{
		"name":       updated.Name,
		"email":      updated.Email,
		"department": updated.Department,
		"salary":     updated.Salary,
	}).Error
	if err != nil {
		return nil, err
	}

	if err := r.db.First(&employee, id).Error; err != nil {
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