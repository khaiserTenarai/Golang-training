package controller

import (
	"employee-management/module"
	"employee-management/service"
)

func AddEmployee(employee module.Employee) error {
	return service.AddEmployee(employee)
}

func GetAllEmployees() []module.Employee {
	return service.GetAllEmployees()
}

func GetEmployeeByID(id int) (module.Employee, error) {
	return service.GetEmployeeByID(id)
}

func DeleteEmployee(id int) error {
	return service.DeleteEmployee(id)
}
