package repository

import (
	"fmt"

	"example.com/employee_management/model"
)

type EmployeeRepository interface {
	AddEmployee(employee model.Employee)
	GetEmployee(id int) (model.Employee, error)
	GetAllEmployees() []model.Employee
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
}

type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

func NewEmployeeRepository() *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{
		employees: []model.Employee{},
	}
}

func (r *EmployeeRepositoryImpl) AddEmployee(employee model.Employee) {
	r.employees = append(r.employees, employee)
}

func (r *EmployeeRepositoryImpl) GetEmployee(id int) (model.Employee, error) {
	for _, employee := range r.employees {
		if employee.ID == id {
			return employee, nil
		}
	}

	return model.Employee{}, fmt.Errorf("employee with ID %d not found", id)
}

func (r *EmployeeRepositoryImpl) GetAllEmployees() []model.Employee {
	return r.employees
}

func (r *EmployeeRepositoryImpl) UpdateEmployee(employee model.Employee) error {
	for i, existingEmployee := range r.employees {
		if existingEmployee.ID == employee.ID {
			r.employees[i] = employee
			return nil
		}
	}

	return fmt.Errorf("employee with ID %d not found", employee.ID)
}

func (r *EmployeeRepositoryImpl) DeleteEmployee(id int) error {
	for i, employee := range r.employees {
		if employee.ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("employee with ID %d not found", id)
}