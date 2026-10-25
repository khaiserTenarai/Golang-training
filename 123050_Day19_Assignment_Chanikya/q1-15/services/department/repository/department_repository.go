package repository

import (
	"errors"

	"employee-management-microservices/services/department/entity"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	DB *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) Create(department *entity.Department) error {
	return r.DB.Create(department).Error
}

func (r *DepartmentRepository) GetAll() ([]entity.Department, error) {
	var departments []entity.Department
	err := r.DB.Order("id").Find(&departments).Error
	return departments, err
}

func (r *DepartmentRepository) GetByID(id uint) (*entity.Department, error) {
	var department entity.Department
	err := r.DB.First(&department, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *DepartmentRepository) Update(department *entity.Department) error {
	return r.DB.Save(department).Error
}

func (r *DepartmentRepository) Delete(id uint) error {
	result := r.DB.Delete(&entity.Department{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
