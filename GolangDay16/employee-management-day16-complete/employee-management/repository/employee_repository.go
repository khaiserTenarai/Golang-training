package repository

import (
	"employee-management/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeRepository interface {
	Create(*model.Employee) error
	GetAll(model.EmployeeFilter) ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}
type employeeRepository struct{ db *gorm.DB }

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository { return &employeeRepository{db} }
func (r *employeeRepository) Create(e *model.Employee) error {
	e.Department = nil
	if err := r.db.Omit(clause.Associations).Create(e).Error; err != nil {
		return err
	}
	// reload so the response contains the department as well
	return r.db.Preload("Department").First(e).Error
}
func (r *employeeRepository) GetAll(f model.EmployeeFilter) ([]model.Employee, error) {
	var e []model.Employee
	q := r.db.Preload("Department").Order("id")
	if f.DepartmentID > 0 {
		q = q.Where("department_id = ?", f.DepartmentID)
	}
	if f.Search != "" {
		like := "%" + f.Search + "%"
		q = q.Where("name ILIKE ? OR email ILIKE ?", like, like)
	}
	if f.Limit > 0 {
		q = q.Limit(f.Limit)
	}
	if f.Offset > 0 {
		q = q.Offset(f.Offset)
	}
	err := q.Find(&e).Error
	return e, err
}
func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var e model.Employee
	if err := r.db.Preload("Department").First(&e, id).Error; err != nil {
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
	x.Salary = e.Salary
	x.DepartmentID = e.DepartmentID
	if err := r.db.Omit(clause.Associations).Save(&x).Error; err != nil {
		return nil, err
	}
	return r.GetByID(id)
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
