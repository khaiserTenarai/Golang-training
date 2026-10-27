package repository

import (
	"context"
	"day19/department-service/model"
	"gorm.io/gorm"
)

// DepartmentRepository keeps SQL/database details out of the service layer.
type DepartmentRepository struct{ db *gorm.DB }

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository { return &DepartmentRepository{db: db} }
func (r *DepartmentRepository) Create(ctx context.Context, d *model.Department) error {
	return r.db.WithContext(ctx).Create(d).Error
}
func (r *DepartmentRepository) List(ctx context.Context) ([]model.Department, error) {
	var items []model.Department
	err := r.db.WithContext(ctx).Order("id").Find(&items).Error
	return items, err
}
func (r *DepartmentRepository) Get(ctx context.Context, id uint) (*model.Department, error) {
	var item model.Department
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *DepartmentRepository) Update(ctx context.Context, item *model.Department) error {
	return r.db.WithContext(ctx).Save(item).Error
}
func (r *DepartmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Department{}, id).Error
}
