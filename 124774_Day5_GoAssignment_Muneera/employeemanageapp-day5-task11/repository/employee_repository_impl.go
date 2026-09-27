package repository

import (
	"errors"

	"employee-management-app/model"
)

type EmployeeRepositoryImpl struct {

	// Slice acts as our database.
	employees []model.Employee
}

// Creates a new Repository.
func NewEmployeeRepository() EmployeeRepository {

	return &EmployeeRepositoryImpl{
		employees: []model.Employee{},
	}
}

// Add employee.
func (r *EmployeeRepositoryImpl) Save(
	employee model.Employee,
) error {

	// Check duplicate ID.
	for _, emp := range r.employees {

		if emp.ID == employee.ID {
			return errors.New("employee ID already exists")
		}
	}

	r.employees = append(r.employees, employee)

	return nil
}

// Delete employee.
func (r *EmployeeRepositoryImpl) Delete(id int) error {

	for i, emp := range r.employees {

		if emp.ID == id {

			r.employees = append(
				r.employees[:i],
				r.employees[i+1:]...,
			)

			return nil
		}
	}

	return errors.New("employee not found")
}

// Update employee.
func (r *EmployeeRepositoryImpl) Update(
	employee model.Employee,
) error {

	for i, emp := range r.employees {

		if emp.ID == employee.ID {

			r.employees[i] = employee

			return nil
		}
	}

	return errors.New("employee not found")
}

// Find employee by ID.
func (r *EmployeeRepositoryImpl) FindByID(
	id int,
) (model.Employee, error) {

	for _, emp := range r.employees {

		if emp.ID == id {
			return emp, nil
		}
	}

	return model.Employee{}, errors.New("employee not found")
}

// Find all employees.
func (r *EmployeeRepositoryImpl) FindAll() []model.Employee {

	return r.employees
}
