package repository

import (
	"errors"

	"employee-management/model"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	Add(emp model.Employee)
	GetByID(id int) (model.Employee, error)
	Delete(id int) error
}

type InMemoryEmployeeRepository struct {
	employees []model.Employee
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{}
}

func (r *InMemoryEmployeeRepository) Add(emp model.Employee) {
	r.employees = append(r.employees, emp)
}

func (r *InMemoryEmployeeRepository) GetByID(id int) (model.Employee, error) {
	for _, emp := range r.employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return model.Employee{}, ErrEmployeeNotFound
}

func (r *InMemoryEmployeeRepository) Delete(id int) error {
	for i, emp := range r.employees {
		if emp.ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}
	return ErrEmployeeNotFound
}
