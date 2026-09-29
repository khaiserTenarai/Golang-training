package repository

import (
	"errors"

	"app/model"
)

type EmployeeRepository interface {
	Send(emp *model.Employee) error        
	Receive(id int) (*model.Employee, error)      
	ReceiveAll() ([]model.Employee, error)          
	Delete(id int) error                             
}

type MemoryEmployeeRepository struct {
	storage map[int]*model.Employee
}

func NewMemoryEmployeeRepository() EmployeeRepository {
	return &MemoryEmployeeRepository{
		storage: make(map[int]*model.Employee),
	}
}

func (r *MemoryEmployeeRepository) Send(emp *model.Employee) error {
	if emp == nil {
		return errors.New("cannot send nil employee")
	}
	r.storage[emp.ID] = emp
	return nil
}

func (r *MemoryEmployeeRepository) Receive(id int) (*model.Employee, error) {
	emp, exists := r.storage[id]
	if !exists {
		return nil, errors.New("employee not found")
	}
	return emp, nil
}

func (r *MemoryEmployeeRepository) ReceiveAll() ([]model.Employee, error) {
	list := make([]model.Employee, 0, len(r.storage))
	for _, emp := range r.storage {
		list = append(list, *emp)
	}
	return list, nil
}

func (r *MemoryEmployeeRepository) Delete(id int) error {
	if _, exists := r.storage[id]; !exists {
		return errors.New("employee not found")
	}
	delete(r.storage, id)
	return nil
}