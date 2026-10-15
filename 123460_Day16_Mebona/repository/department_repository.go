package repository

import (
	"employee-management/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	List() ([]*model.Department, error)
	GetByID(id uint) (*model.Department, error)
}

type departmentRepository struct{ db *gorm.DB }

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository { return &departmentRepository{db} }

func (r *departmentRepository) List() ([]*model.Department, error) {
	var out []*model.Department
	err := r.db.Order("id").Find(&out).Error
	return out, err
}

func (r *departmentRepository) GetByID(id uint) (*model.Department, error) {
	var d model.Department
	if err := r.db.First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}
