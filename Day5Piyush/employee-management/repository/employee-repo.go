package repository

import (
	"errors"

	"Day5Piyush/employee-management/model"
)

var (
	ErrEmployeeNotFound  = errors.New("employee not found")
	ErrDuplicateEmployee = errors.New("employee already exists")
)

type InMemoryEmployeeRepository struct {
	employees map[int]model.Employee
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make(map[int]model.Employee),
	}
}

func (r *InMemoryEmployeeRepository) Add(employee model.Employee) error {

	if _, exists := r.employees[employee.ID]; exists {
		return ErrDuplicateEmployee
	}

	r.employees[employee.ID] = employee

	return nil
}

func (r *InMemoryEmployeeRepository) GetByID(
	id int,
) (model.Employee, error) {

	employee, exists := r.employees[id]

	if !exists {
		return model.Employee{}, ErrEmployeeNotFound
	}

	return employee, nil
}

func (r *InMemoryEmployeeRepository) GetAll() []model.Employee {

	employees := make([]model.Employee, 0, len(r.employees))

	for _, employee := range r.employees {
		employees = append(employees, employee)
	}

	return employees
}

func (r *InMemoryEmployeeRepository) Delete(id int) error {

	if _, exists := r.employees[id]; !exists {
		return ErrEmployeeNotFound
	}

	delete(r.employees, id)

	return nil
}
