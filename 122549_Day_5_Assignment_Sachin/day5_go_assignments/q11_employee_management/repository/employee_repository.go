package repository

import (
	"errors"

	"employee-app/model"
)

type EmployeeRepository interface {
	Save(emp model.Employee) error
	GetByID(id int) (model.Employee, error)
	GetAll() []model.Employee
}

type MemoryEmployeeRepository struct {
	data map[int]model.Employee
}

func NewMemoryEmployeeRepository() *MemoryEmployeeRepository {
	return &MemoryEmployeeRepository{
		data: make(map[int]model.Employee),
	}
}

func (r *MemoryEmployeeRepository) Save(emp model.Employee) error {
	r.data[emp.ID] = emp
	return nil
}

func (r *MemoryEmployeeRepository) GetByID(id int) (model.Employee, error) {
	emp, ok := r.data[id]
	if !ok {
		return model.Employee{}, errors.New("employee not found")
	}
	return emp, nil
}

func (r *MemoryEmployeeRepository) GetAll() []model.Employee {
	list := make([]model.Employee, 0, len(r.data))
	for _, emp := range r.data {
		list = append(list, emp)
	}
	return list
}
