package repository

import (
	"employee-management/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(*model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}
type employeeRepository struct{ db *gorm.DB }

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository   { return &employeeRepository{db} }
func (r *employeeRepository) Create(e *model.Employee) error { return r.db.Create(e).Error }
func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	var e []model.Employee
	err := r.db.Order("id").Find(&e).Error
	return e, err
}
func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var e model.Employee
	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}
func (r *employeeRepository) Update(id uint, e *model.Employee) (*model.Employee, error) {
	var x model.Employee
	if err := r.db.First(&x, id).Error; err != nil {
		return nil, err
	}
	x.Name = e.Name
	x.Email = e.Email
	x.Department = e.Department
	x.Salary = e.Salary
	if err := r.db.Save(&x).Error; err != nil {
		return nil, err
	}
	return &x, nil
}
func (r *employeeRepository) Delete(id uint) error {
	res := r.db.Delete(&model.Employee{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
