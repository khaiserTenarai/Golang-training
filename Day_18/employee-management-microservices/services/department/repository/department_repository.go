package repository

import (
	"errors"

	"employee-management-microservices/services/department/model"

	"gorm.io/gorm"
)

// Repository defines department database operations.
type Repository interface {
	Create(*model.Department) error
	GetAll() ([]model.Department, error)
	GetByID(uint) (*model.Department, error)
	Update(*model.Department) error
	Delete(uint) error
}

type repository struct {
	db *gorm.DB
}

// New creates a department repository.
func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(d *model.Department) error {
	return r.db.Create(d).Error
}

func (r *repository) GetAll() ([]model.Department, error) {
	var departments []model.Department
	err := r.db.Order("id").Find(&departments).Error
	return departments, err
}

func (r *repository) GetByID(id uint) (*model.Department, error) {
	var department model.Department

	err := r.db.First(&department, id).Error
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *repository) Update(d *model.Department) error {
	result := r.db.Model(&model.Department{}).
		Where("id = ?", d.ID).
		Updates(map[string]interface{}{
			"name":        d.Name,
			"description": d.Description,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) Delete(id uint) error {
	result := r.db.Delete(&model.Department{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// IsNotFound checks whether a record was not found.
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
