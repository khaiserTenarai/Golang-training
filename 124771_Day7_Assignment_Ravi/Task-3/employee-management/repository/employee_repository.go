package repository

import (
	"context"

	"employee-management/model"
	"employee-management/view"
)

type EmployeeRepository interface {
	Search(
		ctx context.Context,
		req view.EmployeeSearchRequest,
	) ([]model.Employee, int64, error)
}
