package repository

import (
	"employee-management/model"
	"fmt"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

// Question 9, 10 and 11:
// Pagination + Filtering + Sorting
func (r *EmployeeRepository) GetEmployees(
	page int,
	limit int,
	department string,
	sortBy string,
	sortOrder string,
) ([]model.Employee, error) {

	var employees []model.Employee

	query := r.DB.Model(&model.Employee{})

	// Question 10: Filter by department
	if department != "" {
		query = query.Where("department = ?", department)
	}

	// Question 11: Allow only safe column names
	allowedColumns := map[string]bool{
		"id":     true,
		"name":   true,
		"salary": true,
	}

	if !allowedColumns[sortBy] {
		sortBy = "id"
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	// Question 9: Pagination
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	err := query.
		Offset(offset).
		Limit(limit).
		Find(&employees).Error

	return employees, err
}

func (r *EmployeeRepository) GetByID(id uint) (model.Employee, error) {
	var employee model.Employee

	err := r.DB.First(&employee, id).Error
	return employee, err
}

func (r *EmployeeRepository) Create(
	employee model.Employee,
) (model.Employee, error) {
	err := r.DB.Create(&employee).Error
	return employee, err
}

func (r *EmployeeRepository) Update(
	id uint,
	updated model.Employee,
) (model.Employee, error) {

	var employee model.Employee

	if err := r.DB.First(&employee, id).Error; err != nil {
		return employee, err
	}

	employee.Name = updated.Name
	employee.Email = updated.Email
	employee.Department = updated.Department
	employee.Salary = updated.Salary

	err := r.DB.Save(&employee).Error
	return employee, err
}

func (r *EmployeeRepository) Delete(id uint) error {
	result := r.DB.Delete(&model.Employee{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
