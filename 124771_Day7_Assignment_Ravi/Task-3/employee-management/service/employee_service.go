package service

import (
	"context"

	"employee-management/view"
)

type EmployeeService interface {
	SearchEmployees(
		ctx context.Context,
		req view.EmployeeSearchRequest,
	) (*view.EmployeeSearchResponse, error)
}
