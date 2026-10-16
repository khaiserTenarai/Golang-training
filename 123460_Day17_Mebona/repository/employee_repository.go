package repository

import (
	"employee-management/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeRepository interface {
	Create(*model.Employee) error
	List(model.EmployeeFilter) ([]model.Employee, int64, error)
	ListByDepartment(uint) ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}
type employeeRepository struct{ db *gorm.DB }

// sortColumns maps the public sort name to the real database column.
// Only names in this list are allowed, so nobody can inject SQL through "sortBy".
var sortColumns = map[string]string{
	"id":        "id",
	"name":      "name",
	"email":     "email",
	"salary":    "salary",
	"createdAt": "created_at",
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository { return &employeeRepository{db} }
func (r *employeeRepository) Create(e *model.Employee) error {
	e.Department = nil
	if err := r.db.Omit(clause.Associations).Create(e).Error; err != nil {
		return err
	}
	// reload so the response contains the department as well
	return r.db.Preload("Department").First(e).Error
}

// filtered builds the WHERE part. It is used for both counting and fetching.
func (r *employeeRepository) filtered(f model.EmployeeFilter) *gorm.DB {
	q := r.db.Model(&model.Employee{})
	if f.DepartmentID > 0 {
		q = q.Where("department_id = ?", f.DepartmentID)
	}
	if f.Search != "" {
		like := "%" + f.Search + "%"
		q = q.Where("(name ILIKE ? OR email ILIKE ?)", like, like)
	}
	if f.MinSalary != nil {
		q = q.Where("salary >= ?", *f.MinSalary)
	}
	if f.MaxSalary != nil {
		q = q.Where("salary <= ?", *f.MaxSalary)
	}
	return q
}

// List returns one page of employees and the total number of matches.
//
// OPTIMIZATION: Preload("Department") loads the departments of ALL employees in
// ONE extra query. Without it, GraphQL would run one query per employee (the "N+1 problem").
func (r *employeeRepository) List(f model.EmployeeFilter) ([]model.Employee, int64, error) {
	var total int64
	if err := r.filtered(f).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	column, ok := sortColumns[f.SortBy]
	if !ok {
		column = "id"
	}
	direction := "ASC"
	if f.SortOrder == "desc" {
		direction = "DESC"
	}
	order := column + " " + direction
	if column != "id" {
		order += ", id ASC" // tie-breaker so pages never overlap
	}
	var items []model.Employee
	err := r.filtered(f).
		Preload("Department").
		Order(order).
		Limit(f.PageSize).
		Offset((f.Page - 1) * f.PageSize).
		Find(&items).Error
	return items, total, err
}
func (r *employeeRepository) ListByDepartment(deptID uint) ([]model.Employee, error) {
	var e []model.Employee
	err := r.db.Preload("Department").Where("department_id = ?", deptID).Order("id").Find(&e).Error
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
