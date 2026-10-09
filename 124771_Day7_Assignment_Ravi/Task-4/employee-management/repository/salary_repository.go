package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"employee-management/model"
)

type SalaryRepository interface {
	UpdateSalary(
		ctx context.Context,
		tx pgx.Tx,
		employeeID int64,
		newSalary float64,
	) error

	CreateSalaryHistory(
		ctx context.Context,
		tx pgx.Tx,
		employeeID int64,
		oldSalary float64,
		newSalary float64,
	) (*model.SalaryHistory, error)

	GetSalaryHistory(
		ctx context.Context,
		employeeID int64,
	) ([]model.SalaryHistory, error)
}
