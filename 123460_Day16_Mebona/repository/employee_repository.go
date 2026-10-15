package repository

import (
	"strings"

	"employee-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(e *model.Employee) error
	GetByID(id uint) (*model.Employee, error)
	GetByEmail(email string) (*model.Employee, error)
	List(f model.EmployeeFilter) ([]*model.Employee, error)
	Update(e *model.Employee) error
	Delete(id uint) error
}

type employeeRepository struct{ db *gorm.DB }

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository { return &employeeRepository{db} }

// Omit("Department") stops GORM from trying to save the nested department struct.
func (r *employeeRepository) Create(e *model.Employee) error {
	return r.db.Omit("Department").Create(e).Error
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var e model.Employee
	if err := r.db.Preload("Department").First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *employeeRepository) GetByEmail(email string) (*model.Employee, error) {
	var e model.Employee
	if err := r.db.Preload("Department").Where("LOWER(email) = LOWER(?)", email).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *employeeRepository) List(f model.EmployeeFilter) ([]*model.Employee, error) {
	q := r.db.Preload("Department").Order("id")
	if f.DepartmentID != nil {
		q = q.Where("department_id = ?", *f.DepartmentID)
	}
	if f.Search != "" {
		like := "%" + strings.ToLower(f.Search) + "%"
		q = q.Where("(LOWER(name) LIKE ? OR LOWER(email) LIKE ?)", like, like)
	}
	if f.JoinedAfter != nil {
		q = q.Where("joined_at > ?", *f.JoinedAfter)
	}
	if f.Limit > 0 {
		q = q.Limit(f.Limit)
	}
	if f.Offset > 0 {
		q = q.Offset(f.Offset)
	}
	var out []*model.Employee
	err := q.Find(&out).Error
	return out, err
}

func (r *employeeRepository) Update(e *model.Employee) error {
	return r.db.Omit("Department").Save(e).Error
}

func (r *employeeRepository) Delete(id uint) error {
	res := r.db.Delete(&model.Employee{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
