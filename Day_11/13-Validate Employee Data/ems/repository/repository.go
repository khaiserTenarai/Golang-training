package repository

import (
	"errors"

	"employee-management-go/model"
)

// EmployeeRepository defines employee operations.
type EmployeeRepository interface {
	Create(employee model.Employee) (model.Employee, error)
	GetAll() []model.Employee
	GetByID(id int) (model.Employee, error)
	Delete(id int) error
}

// EmployeeRepositoryImpl implements EmployeeRepository.
type EmployeeRepositoryImpl struct {
	// Slice stores employees.
	employees []model.Employee
}

// Create adds an employee.
func (r *EmployeeRepositoryImpl) Create(employee model.Employee) (model.Employee, error) {

	employee.ID = len(r.employees) + 1

	r.employees = append(r.employees, employee)

	return employee, nil
}

// GetAll returns all employees.
func (r *EmployeeRepositoryImpl) GetAll() []model.Employee {
	return r.employees
}

// GetByID returns an employee by ID.
func (r *EmployeeRepositoryImpl) GetByID(id int) (model.Employee, error) {

	for _, employee := range r.employees {

		if employee.ID == id {
			return employee, nil
		}
	}

	return model.Employee{}, errors.New("employee not found")
}

// Delete removes an employee.
func (r *EmployeeRepositoryImpl) Delete(id int) error {

	for i, employee := range r.employees {

		if employee.ID == id {

			// Remove employee from slice.
			r.employees = append(
				r.employees[:i],
				r.employees[i+1:]...,
			)

			return nil
		}
	}

	return errors.New("employee not found")
}
