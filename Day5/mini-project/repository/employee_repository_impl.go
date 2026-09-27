package repository

import (
	"fmt"

	"example.com/employee-management/model"
	"example.com/employee-management/util"
)

type EmployeeRepositoryImpl struct {
	employees []model.Employee
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{
		employees: make([]model.Employee, 0),
	}
}

func (r *EmployeeRepositoryImpl) Save(employee model.Employee) model.Employee {
	fmt.Println("Hello from Repository - Save Employee")
	if employee.ID == 0 {
		employee.ID = util.NextID()
	}
	r.employees = append(r.employees, employee)
	return employee
}

func (r *EmployeeRepositoryImpl) FindByID(id int) (model.Employee, bool) {
	fmt.Println("Hello from Repository - Find Employee")
	for _, e := range r.employees {
		if e.ID == id {
			return e, true
		}
	}
	return model.Employee{}, false
}

func (r *EmployeeRepositoryImpl) FindAll() []model.Employee {
	fmt.Println("Hello from Repository - Find All Employees")
	// return a copy so callers can't mutate our internal slice directly
	result := make([]model.Employee, len(r.employees))
	copy(result, r.employees)
	return result
}

func (r *EmployeeRepositoryImpl) Update(employee model.Employee) (model.Employee, bool) {
	fmt.Println("Hello from Repository - Update Employee")
	for i, e := range r.employees {
		if e.ID == employee.ID {
			r.employees[i] = employee
			return employee, true
		}
	}
	return model.Employee{}, false
}

func (r *EmployeeRepositoryImpl) Delete(id int) bool {
	fmt.Println("Hello from Repository - Delete Employee")
	for i, e := range r.employees {
		if e.ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return true
		}
	}
	return false
}
