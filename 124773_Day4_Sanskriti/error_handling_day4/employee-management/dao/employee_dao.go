package dao

import (
	"employee-management/model"
	"employee-management/utility"
)

var employees []model.Employee

func AddEmployee(employee model.Employee) error {

	// Check if employee already exists
	for _, e := range employees {
		if e.ID == employee.ID {
			return utility.ErrDuplicateEmployee
		}
	}

	// Add employee to slice
	employees = append(employees, employee)

	return nil
}

func GetEmployee(id int) (model.Employee, error) {

	for _, employee := range employees {
		if employee.ID == id {
			return employee, nil
		}
	}

	return model.Employee{}, utility.ErrEmployeeNotFound
}

func DeleteEmployee(id int) error {

	for i, employee := range employees {

		if employee.ID == id {

			employees = append(
				employees[:i],
				employees[i+1:]...,
			)

			return nil
		}
	}

	return utility.ErrEmployeeNotFound
}
