package repository

import (
	"employee-management/model"
	"errors"
)

type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{
		employees: make([]model.Employee, 0),
	}
}

func (r *EmployeeRepositoryImpl) Save(employee model.Employee) error {
	r.employees = append(r.employees, employee)
	return nil
}

func (r *EmployeeRepositoryImpl) FindById(id int) (model.Employee, error) {
	for _, employee := range r.employees {
		if employee.ID == id {
			return employee, nil
		}
	}
	return model.Employee{}, errors.New("Employee not found")
}

func (r *EmployeeRepositoryImpl) FindAll() []model.Employee {
	return r.employees
}

func (r *EmployeeRepositoryImpl) Update(employee model.Employee) error {
	for i, existingEmp := range r.employees {
		if existingEmp.ID == employee.ID {
			r.employees[i] = employee
			return nil
		}
	}
	return errors.New("employee not found")
}

func (r *EmployeeRepositoryImpl) Delete(id int) error {
	for i, employee := range r.employees {
		if employee.ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}
	return errors.New("Employee not found")
}
