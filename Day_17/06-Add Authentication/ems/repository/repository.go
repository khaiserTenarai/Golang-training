package repository

import (
	"errors"
	"fmt"
	"strings"

	"employee-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(*model.Employee) error
	GetAll(model.EmployeeFilter, model.EmployeeSort, int, int) ([]model.Employee, int64, error)
	GetByID(uint) (*model.Employee, error)
	Update(uint, *model.Employee) (*model.Employee, error)
	Delete(uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(e *model.Employee) error {
	return r.db.Create(e).Error
}

func (r *employeeRepository) GetAll(
	filter model.EmployeeFilter,
	sort model.EmployeeSort,
	page, pageSize int,
) ([]model.Employee, int64, error) {
	// Apply only parameterized filters to prevent SQL injection.
	query := r.db.Model(&model.Employee{})

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Department != "" {
		query = query.Where("department = ?", filter.Department)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Allowlist the sortable columns. Never concatenate client-provided SQL.
	columns := map[string]string{
		"id": "id", "name": "name", "email": "email",
		"department": "department", "salary": "salary",
	}
	column, ok := columns[strings.ToLower(sort.Field)]
	if !ok {
		column = "id"
	}

	direction := "ASC"
	if strings.EqualFold(sort.Direction, "desc") {
		direction = "DESC"
	}

	var employees []model.Employee
	err := query.
		Order(fmt.Sprintf("%s %s", column, direction)).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&employees).Error

	return employees, total, err
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) Update(id uint, input *model.Employee) (*model.Employee, error) {
	var employee model.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	// Use a map so salary=0 is a valid update.
	err := r.db.Model(&employee).Updates(map[string]interface{}{
		"name": input.Name, "email": input.Email,
		"department": input.Department, "salary": input.Salary,
	}).Error
	if err != nil {
		return nil, err
	}

	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	return nil
}
