package repository

import (
	"errors"
	"strings"

	"day17-employee-platform/model"

	"gorm.io/gorm"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	GetAll(filter model.EmployeeFilter, sort model.SortOption, page, pageSize int) ([]model.Employee, int64, error)
	GetByID(id uint) (*model.Employee, error)
	Create(employee *model.Employee) error
	Update(id uint, employee *model.Employee) (*model.Employee, error)
	Delete(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) GetAll(
	filter model.EmployeeFilter,
	sort model.SortOption,
	page, pageSize int,
) ([]model.Employee, int64, error) {

	query := r.db.Model(&model.Employee{}).Preload("Department")

	if strings.TrimSpace(filter.Name) != "" {
		query = query.Where("name ILIKE ?", "%"+strings.TrimSpace(filter.Name)+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSortFields := map[string]string{
		"id":     "id",
		"name":   "name",
		"salary": "salary",
	}

	sortField := allowedSortFields[sort.Field]
	if sortField == "" {
		sortField = "id"
	}

	sortOrder := "asc"
	if strings.ToLower(sort.Order) == "desc" {
		sortOrder = "desc"
	}

	offset := (page - 1) * pageSize

	var employees []model.Employee
	err := query.
		Order(sortField + " " + sortOrder).
		Offset(offset).
		Limit(pageSize).
		Find(&employees).Error

	return employees, total, err
}

func (r *employeeRepository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee

	err := r.db.Preload("Department").First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrEmployeeNotFound
	}
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) Create(employee *model.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) Update(id uint, employee *model.Employee) (*model.Employee, error) {
	var existing model.Employee

	if err := r.db.First(&existing, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrEmployeeNotFound
	} else if err != nil {
		return nil, err
	}

	existing.Name = employee.Name
	existing.Email = employee.Email
	existing.Salary = employee.Salary
	existing.DepartmentID = employee.DepartmentID

	if err := r.db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *employeeRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrEmployeeNotFound
	}

	return nil
}
