package repository

import (
	"employee-management/model"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(*model.Department) error
	GetAll() ([]model.Department, error)
	GetByID(uint) (*model.Department, error)
}
type departmentRepository struct{ db *gorm.DB }

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db}
}
func (r *departmentRepository) Create(d *model.Department) error { return r.db.Create(d).Error }
func (r *departmentRepository) GetAll() ([]model.Department, error) {
	var d []model.Department
	err := r.db.Order("id").Find(&d).Error
	return d, err
}
func (r *departmentRepository) GetByID(id uint) (*model.Department, error) {
	var d model.Department
	if err := r.db.First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}
