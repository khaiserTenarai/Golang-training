package repository

import (
	"day17-employee-platform/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	GetByID(id uint) (*model.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) GetByID(id uint) (*model.Department, error) {
	var department model.Department
	if err := r.db.First(&department, id).Error; err != nil {
		return nil, err
	}
	return &department, nil
}
