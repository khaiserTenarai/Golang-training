package controller
import "employee-management/model"

type EmployeeController interface {
	AddEmployee(employee model.Employee)
	GetAllEmployees()
	GetEmployeeByID(id int)
	UpdateEmployee(employee model.Employee)
	DeleteEmployee(id int)
}