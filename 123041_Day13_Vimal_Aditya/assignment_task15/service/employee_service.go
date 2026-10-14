package service

import (
	"context"

	"employeeapi/model"
)

type EmployeeService interface {
	AddEmployee(ctx context.Context, employee model.Employee) (model.Employee, error)
	GetEmployee(ctx context.Context, id int) (model.Employee, error)
	GetAllEmployees(ctx context.Context) ([]model.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
	SalaryReport(ctx context.Context) (model.SalaryReport, error)
}
