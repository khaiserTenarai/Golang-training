package repository

import (
	"context"
	"day19/employee-service/model"
	"gorm.io/gorm"
)

// EmployeeRepository contains only persistence operations.
type EmployeeRepository struct{ db *gorm.DB }

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository { return &EmployeeRepository{db: db} }
func (r *EmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	return r.db.WithContext(ctx).Create(e).Error
}
func (r *EmployeeRepository) List(ctx context.Context) ([]model.Employee, error) {
	var rows []model.Employee
	err := r.db.WithContext(ctx).Order("id").Find(&rows).Error
	return rows, err
}
func (r *EmployeeRepository) Get(ctx context.Context, id uint) (*model.Employee, error) {
	var e model.Employee
	err := r.db.WithContext(ctx).First(&e, id).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func (r *EmployeeRepository) Update(ctx context.Context, e *model.Employee) error {
	return r.db.WithContext(ctx).Save(e).Error
}
func (r *EmployeeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Employee{}, id).Error
}
