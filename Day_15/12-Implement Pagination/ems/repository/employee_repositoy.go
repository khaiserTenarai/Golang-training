package repository

import (
	"ems/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(*model.Employee) error
	GetAll(page, pageSize int) ([]model.Employee, int64, error)
	GetByID(uint) (*model.Employee, error)
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
	return r.db.Create(e).Error
}

func (r *employeeRepository) GetAll(
	page, pageSize int,
) ([]model.Employee, int64, error) {
	var employees []model.Employee
	var total int64

	// Count all matching records before pagination.
	if err := r.db.Model(&model.Employee{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	err := r.db.Order("id ASC").
		Limit(pageSize).
		Offset(offset).
		Find(&employees).Error

	return employees, total, err
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var e model.Employee

	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *employeeRepository) Update(
	id uint,
	updated *model.Employee,
) (*model.Employee, error) {
	var e model.Employee

	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}

	err := r.db.Model(&e).Updates(map[string]interface{}{
		"name":       updated.Name,
		"email":      updated.Email,
		"department": updated.Department,
		"salary":     updated.Salary,
	}).Error

	if err != nil {
		return nil, err
	}

	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}

	return &e, nil
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
