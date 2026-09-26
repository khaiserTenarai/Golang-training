package service

import (
	"employee-management/module"
	"employee-management/repository"
	"employee-management/utility"
)

func AddEmployee(employee module.Employee) error {

	if !utility.IsValidName(employee.Name) {
		return utility.ErrInvalidName
	}

	if !utility.IsValidAge(employee.Age) {
		return utility.ErrInvalidAge
	}

	if !utility.IsValidSalary(employee.Salary) {
		return utility.ErrInvalidSalary
	}

	if !utility.IsValidPosition(employee.Position) {
		return utility.ErrInvalidPosition
	}

	return repository.AddEmployee(employee)
}

func GetAllEmployees() []module.Employee {
	return repository.GetAllEmployees()
}

func GetEmployeeByID(id int) (module.Employee, error) {

	if id <= 0 {
		return module.Employee{}, utility.ErrEmployeeNotFound
	}

	return repository.GetEmployeeByID(id)
}

func DeleteEmployee(id int) error {

	if id <= 0 {
		return utility.ErrEmployeeNotFound
	}

	return repository.DeleteEmployee(id)
}
