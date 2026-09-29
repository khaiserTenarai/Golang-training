package repository

import "employee-management/model"

type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{
		employees: make([]model.Employee, 0),
	}
}

func (r *EmployeeRepositoryImpl) Add(employee model.Employee) {
	r.employees = append(r.employees, employee)
}

func (r *EmployeeRepositoryImpl) GetByID(id int) (model.Employee, bool) {
	for _, employee := range r.employees {
		if employee.ID == id {
			return employee, true
		}
	}

	return model.Employee{}, false
}

func (r *EmployeeRepositoryImpl) GetAll() []model.Employee {
	return r.employees
}

func (r *EmployeeRepositoryImpl) Delete(id int) {
	for i, employee := range r.employees {
		if employee.ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return
		}
	}
}
