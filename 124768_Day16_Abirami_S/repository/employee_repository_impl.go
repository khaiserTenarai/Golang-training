package repository

import (
	"errors"
	"sort"

	"day16-employee-graphql/database"
	"day16-employee-graphql/model"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type employeeRepository struct {
	db *database.DB
}

func NewEmployeeRepository(db *database.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) GetByID(id int) (*model.Employee, error) {
	r.db.Mu.RLock()
	defer r.db.Mu.RUnlock()

	employee, ok := r.db.Employees[id]
	if !ok {
		return nil, ErrEmployeeNotFound
	}
	copy := employee
	return &copy, nil
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	r.db.Mu.RLock()
	defer r.db.Mu.RUnlock()

	employees := make([]model.Employee, 0, len(r.db.Employees))
	for _, employee := range r.db.Employees {
		employees = append(employees, employee)
	}

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].ID < employees[j].ID
	})
	return employees, nil
}

func (r *employeeRepository) Create(employee model.Employee) (*model.Employee, error) {
	r.db.Mu.Lock()
	defer r.db.Mu.Unlock()

	employee.ID = r.db.NextID
	r.db.NextID++
	r.db.Employees[employee.ID] = employee

	copy := employee
	return &copy, nil
}

func (r *employeeRepository) Update(id int, employee model.Employee) (*model.Employee, error) {
	r.db.Mu.Lock()
	defer r.db.Mu.Unlock()

	if _, ok := r.db.Employees[id]; !ok {
		return nil, ErrEmployeeNotFound
	}

	employee.ID = id
	r.db.Employees[id] = employee

	copy := employee
	return &copy, nil
}

func (r *employeeRepository) Delete(id int) (*model.Employee, error) {
	r.db.Mu.Lock()
	defer r.db.Mu.Unlock()

	employee, ok := r.db.Employees[id]
	if !ok {
		return nil, ErrEmployeeNotFound
	}

	delete(r.db.Employees, id)
	copy := employee
	return &copy, nil
}
