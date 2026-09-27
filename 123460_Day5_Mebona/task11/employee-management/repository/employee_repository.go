package repository

import "employee-management/model"

type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

func (r *EmployeeRepositoryImpl) Add(employee model.Employee) {
	r.employees = append(r.employees, employee)
}

func (r *EmployeeRepositoryImpl) Get(id int) (model.Employee, bool) {
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

func (r *EmployeeRepositoryImpl) Update(employee model.Employee) bool {
	for i := range r.employees {
		if r.employees[i].ID == employee.ID {
			r.employees[i] = employee
			return true
		}
	}

	return false
}

func (r *EmployeeRepositoryImpl) Delete(id int) bool {
	for i, employee := range r.employees {
		if employee.ID == id {
			r.employees = append(
				r.employees[:i],
				r.employees[i+1:]...,
			)
			return true
		}
	}

	return false
}