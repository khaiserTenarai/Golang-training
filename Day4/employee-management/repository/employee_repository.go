package repository

import (
	"employee-management/module"
	"employee-management/utility"
)

var employees []module.Employee

func AddEmployee(employee module.Employee) error {

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == employee.ID {
			return utility.ErrEmployeeExists
		}
	}

	employees = append(employees, employee)

	return nil
}

func GetAllEmployees() []module.Employee {
	return employees
}

func GetEmployeeByID(id int) (module.Employee, error) {

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			return employees[i], nil
		}
	}

	return module.Employee{}, utility.ErrEmployeeNotFound
}

func DeleteEmployee(id int) error {

	for i := 0; i < len(employees); i++ {

		if employees[i].ID == id {

			employees = append(
				employees[:i],
				employees[i+1:]...,
			)

			return nil
		}
	}

	return utility.ErrEmployeeNotFound
}
