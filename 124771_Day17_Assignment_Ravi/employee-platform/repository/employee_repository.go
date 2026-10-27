package repository

import (
	"employee-platform/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct{ db *gorm.DB }

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository { return &EmployeeRepository{db: db} }
func (r *EmployeeRepository) List(page, limit int, departmentID uint, search, sortBy, order string) ([]model.Employee, int64, error) {
	var rows []model.Employee
	var total int64
	q := r.db.Model(&model.Employee{})
	if departmentID > 0 {
		q = q.Where("department_id = ?", departmentID)
	}
	if search != "" {
		q = q.Where("LOWER(name) LIKE LOWER(?) OR LOWER(email) LIKE LOWER(?)", "%"+search+"%", "%"+search+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Task 9/10/11: pagination, filtering and sorting. Whitelist sort columns to prevent SQL injection.
	allowed := map[string]bool{"id": true, "name": true, "email": true, "salary": true, "created_at": true}
	if !allowed[sortBy] {
		sortBy = "id"
	}
	if order != "desc" {
		order = "asc"
	}
	err := q.Preload("Department").Order(sortBy + " " + order).
		Offset((page - 1) * limit).Limit(limit).Find(&rows).Error
	return rows, total, err
}
func (r *EmployeeRepository) ByID(id uint) (*model.Employee, error) {
	var e model.Employee
	err := r.db.Preload("Department").First(&e, id).Error
	return &e, err
}
func (r *EmployeeRepository) Create(e *model.Employee) error { return r.db.Create(e).Error }
func (r *EmployeeRepository) Update(e *model.Employee) error { return r.db.Save(e).Error }
func (r *EmployeeRepository) Delete(id uint) error           { return r.db.Delete(&model.Employee{}, id).Error }
