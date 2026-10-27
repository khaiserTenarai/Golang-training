package repository

import (
	"employee-platform/model"
	"gorm.io/gorm"
)

type DepartmentRepository struct{ db *gorm.DB }

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository { return &DepartmentRepository{db: db} }
func (r *DepartmentRepository) List() ([]model.Department, error) {
	var rows []model.Department
	err := r.db.Order("id asc").Find(&rows).Error
	return rows, err
}
func (r *DepartmentRepository) ByID(id uint) (*model.Department, error) {
	var d model.Department
	err := r.db.First(&d, id).Error
	return &d, err
}
func (r *DepartmentRepository) Create(d *model.Department) error { return r.db.Create(d).Error }
func (r *DepartmentRepository) Exists(id uint) bool {
	var count int64
	return r.db.Model(&model.Department{}).Where("id = ?", id).Count(&count).Error == nil && count > 0
}
