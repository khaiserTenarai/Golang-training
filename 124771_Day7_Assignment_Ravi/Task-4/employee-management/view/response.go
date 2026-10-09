package view

import "employee-management/model"

type SalaryUpdateResponse struct {
	Employee      model.Employee
	SalaryHistory model.SalaryHistory
}

type SalaryHistoryResponse struct {
	EmployeeID int64
	History    []model.SalaryHistory
}
