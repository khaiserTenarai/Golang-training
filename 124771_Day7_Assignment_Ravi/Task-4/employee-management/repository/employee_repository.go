package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"employee-management/model"
)

type EmployeeRepository interface {
	GetByID(
		ctx context.Context,
		tx pgx.Tx,
		employeeID int64,
	) (*model.Employee, error)
}
