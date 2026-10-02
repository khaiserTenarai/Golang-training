package service

import "example.com/employee-management/model"

type EmployeeService interface {
    AddEmployee(employee model.Employee) error
    GetEmployee(id int64) (model.Employee, error)
    GetAllEmployees() ([]model.Employee, error)
    UpdateEmployee(employee model.Employee) error
    DeleteEmployee(id int64) error
}
