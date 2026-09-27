package repository

import "example.com/employee-management/model"

// EmployeeRepositoryImpl keeps all employees in memory using a slice.
// This is our "database" for this simple version - nothing is written
// to disk, so data only lasts for as long as the program is running.
type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{
		employees: []model.Employee{},
	}
}

func (r *EmployeeRepositoryImpl) Save(employee model.Employee) {
	r.employees = append(r.employees, employee)
}

func (r *EmployeeRepositoryImpl) FindAll() []model.Employee {
	return r.employees
}
