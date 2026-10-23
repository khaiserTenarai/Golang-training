package repository

import (
	"employee-management-microservices/events/outbox"
	"employee-management-microservices/services/employee/model"

	"gorm.io/gorm"
)

// Repository defines employee database operations.
type Repository interface {
	Create(*model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(uint) (*model.Employee, error)
	Update(*model.Employee) error
	Delete(uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create saves the employee and event in one transaction.
func (r *repository) Create(e *model.Employee) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(e).Error; err != nil {
			return err
		}

		event, err := outbox.NewEvent("EmployeeCreated", e)
		if err != nil {
			return err
		}

		return outbox.Save(tx, event)
	})
}

func (r *repository) GetAll() ([]model.Employee, error) {
	var employees []model.Employee
	err := r.db.Order("id").Find(&employees).Error
	return employees, err
}

func (r *repository) GetByID(id uint) (*model.Employee, error) {
	var employee model.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

// Update saves the employee and its event together.
func (r *repository) Update(e *model.Employee) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Employee{}).
			Where("id = ?", e.ID).
			Updates(map[string]interface{}{
				"name":          e.Name,
				"email":         e.Email,
				"department_id": e.DepartmentID,
				"salary":        e.Salary,
			})

		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		event, err := outbox.NewEvent("EmployeeUpdated", e)
		if err != nil {
			return err
		}

		return outbox.Save(tx, event)
	})
}

// Delete saves the event before deleting the employee.
func (r *repository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var employee model.Employee

		if err := tx.First(&employee, id).Error; err != nil {
			return err
		}

		event, err := outbox.NewEvent("EmployeeDeleted", employee)
		if err != nil {
			return err
		}

		if err := outbox.Save(tx, event); err != nil {
			return err
		}

		return tx.Delete(&employee).Error
	})
}
