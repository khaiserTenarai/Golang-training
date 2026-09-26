package repository

import "employee-management/model"

// EmployeeRepository defines employee operations
type EmployeeRepository interface {
	AddEmployee(employee model.Employee)
	GetEmployee(id int) (model.Employee, bool)
	DeleteEmployee(id int) bool
	GetAllEmployees() []model.Employee
}

// EmployeeRepositoryImpl stores employees in memory
type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

// Create a new repository
func NewEmployeeRepository() *EmployeeRepositoryImpl {

	return &EmployeeRepositoryImpl{
		employees: []model.Employee{},
	}
}

// Add employee
func (r *EmployeeRepositoryImpl) AddEmployee(employee model.Employee) {

	r.employees = append(r.employees, employee)
}

// Get employee
func (r *EmployeeRepositoryImpl) GetEmployee(id int) (model.Employee, bool) {

	for _, employee := range r.employees {

		if employee.ID == id {
			return employee, true
		}
	}

	return model.Employee{}, false
}

// Delete employee
func (r *EmployeeRepositoryImpl) DeleteEmployee(id int) bool {

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

// Get all employees
func (r *EmployeeRepositoryImpl) GetAllEmployees() []model.Employee {

	return r.employees
}