package repository

import "employee-management/model"


// Dependency Injection
type EmployeeRepositoryImpl struct {
	employees []model.Employee
}


// func retu

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{
		employees: make([]model.Employee, 0),
	}
}

func (r *EmployeeRepositoryImpl) AddEmployee(employee model.Employee) error {
	r.employees = append(r.employees, employee)
	return nil
}

func (r *EmployeeRepositoryImpl) GetAllEmployees() []model.Employee {
	return r.employees
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(id int) model.Employee {

	for _, employee := range r.employees {
		if employee.ID == id {
			return employee
		}
	}

	return model.Employee{}
}

func (r *EmployeeRepositoryImpl) UpdateEmployee(
	employee model.Employee,
) error {

	for i := range r.employees {

		if r.employees[i].ID == employee.ID {
			r.employees[i] = employee
			return nil
		}
	}

	return nil
}

func (r *EmployeeRepositoryImpl) DeleteEmployee(id int) error {

	for i := range r.employees {

		if r.employees[i].ID == id {
			r.employees = append(
				r.employees[:i],
				r.employees[i+1:]...,
			)

			return nil
		}
	}

	return nil
}