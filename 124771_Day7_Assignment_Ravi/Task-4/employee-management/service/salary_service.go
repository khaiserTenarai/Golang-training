package service

import (
	"context"

	"employee-management/view"
)

type SalaryService interface {
	UpdateSalary(
		ctx context.Context,
		req view.SalaryUpdateRequest,
	) (*view.SalaryUpdateResponse, error)

	GetSalaryHistory(
		ctx context.Context,
		employeeID int64,
	) (*view.SalaryHistoryResponse, error)
}
