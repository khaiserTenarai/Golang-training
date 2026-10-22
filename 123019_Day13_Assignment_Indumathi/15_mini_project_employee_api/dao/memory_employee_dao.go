package dao

import (
	"context"
	"sync"

	"employeeapi/model"
)

type MemoryEmployeeDAO struct {
	mu     sync.Mutex
	items  map[int]model.Employee
	nextID int
}

func NewMemoryEmployeeDAO() *MemoryEmployeeDAO {
	return &MemoryEmployeeDAO{items: map[int]model.Employee{}, nextID: 1}
}

func (d *MemoryEmployeeDAO) Save(ctx context.Context, employee model.Employee) (model.Employee, error) {
	if err := ctx.Err(); err != nil {
		return model.Employee{}, err
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	employee.ID = d.nextID
	d.nextID++
	d.items[employee.ID] = employee
	return employee, nil
}

func (d *MemoryEmployeeDAO) FindByID(ctx context.Context, id int) (model.Employee, error) {
	if err := ctx.Err(); err != nil {
		return model.Employee{}, err
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	employee, ok := d.items[id]
	if !ok {
		return model.Employee{}, ErrNotFound
	}
	return employee, nil
}

func (d *MemoryEmployeeDAO) FindAll(ctx context.Context) ([]model.Employee, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	list := make([]model.Employee, 0, len(d.items))
	for id := 1; id < d.nextID; id++ {
		if employee, ok := d.items[id]; ok {
			list = append(list, employee)
		}
	}
	return list, nil
}

func (d *MemoryEmployeeDAO) Delete(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.items[id]; !ok {
		return ErrNotFound
	}
	delete(d.items, id)
	return nil
}
