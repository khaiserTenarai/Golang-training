package view

import "employee-management/model"

type EmployeeSearchResponse struct {
	Data       []model.Employee
	Page       int
	PageSize   int
	Total      int64
	TotalPages int
}
