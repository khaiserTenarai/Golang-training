package dao

import (
	"strings"
	"sync"

	"employeegin/model"
)

type MemoryEmployeeDAO struct {
	mu     sync.RWMutex
	items  map[int]model.Employee
	nextID int
}

func NewMemoryEmployeeDAO() *MemoryEmployeeDAO {
	return &MemoryEmployeeDAO{items: map[int]model.Employee{}, nextID: 1}
}

func (d *MemoryEmployeeDAO) Save(employee model.Employee) (model.Employee, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	employee.ID = d.nextID
	d.nextID++
	d.items[employee.ID] = employee
	return employee, nil
}

func (d *MemoryEmployeeDAO) FindByID(id int) (model.Employee, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	employee, ok := d.items[id]
	if !ok {
		return model.Employee{}, ErrNotFound
	}
	return employee, nil
}

func (d *MemoryEmployeeDAO) FindByEmail(email string) (model.Employee, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, employee := range d.items {
		if strings.EqualFold(employee.Email, email) {
			return employee, nil
		}
	}
	return model.Employee{}, ErrNotFound
}

func (d *MemoryEmployeeDAO) FindAll(department string) ([]model.Employee, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	list := []model.Employee{}
	for id := 1; id < d.nextID; id++ {
		employee, ok := d.items[id]
		if !ok {
			continue
		}
		if department != "" && !strings.EqualFold(employee.Department, department) {
			continue
		}
		list = append(list, employee)
	}
	return list, nil
}

func (d *MemoryEmployeeDAO) Update(employee model.Employee) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.items[employee.ID]; !ok {
		return ErrNotFound
	}
	d.items[employee.ID] = employee
	return nil
}

func (d *MemoryEmployeeDAO) Delete(id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.items[id]; !ok {
		return ErrNotFound
	}
	delete(d.items, id)
	return nil
}
