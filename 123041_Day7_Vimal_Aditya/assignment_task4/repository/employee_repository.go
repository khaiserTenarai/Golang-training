package repository

import "example.com/employee-management/model"

type EmployeeRepository interface {
    Save(employee model.Employee) error
    FindByID(id int64) (model.Employee, error)
    FindAll() ([]model.Employee, error)
    Update(employee model.Employee) error
    Delete(id int64) error
    UpdateSalaryWithHistory(id int64, newSalary float64) error
    GetSalaryHistory(employeeID int64) ([]model.SalaryHistory, error)
}