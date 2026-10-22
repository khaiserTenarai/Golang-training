package repository

import (
	"errors"

	"employee-management-microservices/services/employee/entity"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *entity.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *EmployeeRepository) GetAll() ([]entity.Employee, error) {
	var employees []entity.Employee
	err := r.DB.Order("id").Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepository) GetByID(id uint) (*entity.Employee, error) {
	var employee entity.Employee
	err := r.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) Update(employee *entity.Employee) error {
	return r.DB.Save(employee).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	result := r.DB.Delete(&entity.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
