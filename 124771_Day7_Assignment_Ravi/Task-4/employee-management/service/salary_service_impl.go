package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/repository"
	"employee-management/utility"
	"employee-management/view"
)

type SalaryServiceImpl struct {
	db                 *pgxpool.Pool
	employeeRepository repository.EmployeeRepository
	salaryRepository   repository.SalaryRepository
}

func NewSalaryService(
	db *pgxpool.Pool,
	employeeRepository repository.EmployeeRepository,
	salaryRepository repository.SalaryRepository,
) *SalaryServiceImpl {

	return &SalaryServiceImpl{
		db:                 db,
		employeeRepository: employeeRepository,
		salaryRepository:   salaryRepository,
	}
}

// ----------------------------------------------------
// Update Salary
// ----------------------------------------------------

func (s *SalaryServiceImpl) UpdateSalary(
	ctx context.Context,
	req view.SalaryUpdateRequest,
) (*view.SalaryUpdateResponse, error) {

	// --------------------------------------------
	// Validation
	// --------------------------------------------

	if err := utility.ValidateEmployeeID(
		req.EmployeeID,
	); err != nil {

		return nil, err
	}

	if err := utility.ValidateSalary(
		req.NewSalary,
	); err != nil {

		return nil, err
	}

	// --------------------------------------------
	// Begin Transaction
	// --------------------------------------------

	tx, err := s.db.BeginTx(
		ctx,
		pgx.TxOptions{},
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to begin transaction: %w",
			err,
		)
	}

	// --------------------------------------------
	// Rollback automatically if something fails
	// --------------------------------------------

	committed := false

	defer func() {

		if !committed {
			_ = tx.Rollback(ctx)
		}

	}()

	// --------------------------------------------
	// Get Employee FOR UPDATE
	// --------------------------------------------

	employee, err :=
		s.employeeRepository.GetByID(
			ctx,
			tx,
			req.EmployeeID,
		)

	if err != nil {

		if errors.Is(
			err,
			repository.ErrEmployeeNotFound,
		) {
			return nil, err
		}

		return nil, fmt.Errorf(
			"failed to get employee: %w",
			err,
		)
	}

	// --------------------------------------------
	// Check same salary
	// --------------------------------------------

	if employee.Salary == req.NewSalary {

		return nil, fmt.Errorf(
			"new salary is same as current salary",
		)
	}

	oldSalary := employee.Salary

	// --------------------------------------------
	// Update Employee Salary
	// --------------------------------------------

	err = s.salaryRepository.UpdateSalary(
		ctx,
		tx,
		req.EmployeeID,
		req.NewSalary,
	)

	if err != nil {

		return nil, fmt.Errorf(
			"failed to update salary: %w",
			err,
		)
	}

	// --------------------------------------------
	// Insert Salary History
	// --------------------------------------------

	history, err :=
		s.salaryRepository.CreateSalaryHistory(
			ctx,
			tx,
			req.EmployeeID,
			oldSalary,
			req.NewSalary,
		)

	if err != nil {

		return nil, fmt.Errorf(
			"failed to create salary history: %w",
			err,
		)
	}

	// --------------------------------------------
	// Commit Transaction
	// --------------------------------------------

	if err := tx.Commit(ctx); err != nil {

		return nil, fmt.Errorf(
			"failed to commit transaction: %w",
			err,
		)
	}

	committed = true

	// Update returned employee object
	employee.Salary = req.NewSalary

	return &view.SalaryUpdateResponse{
		Employee:      *employee,
		SalaryHistory: *history,
	}, nil
}

// ----------------------------------------------------
// Get Salary History
// ----------------------------------------------------

func (s *SalaryServiceImpl) GetSalaryHistory(
	ctx context.Context,
	employeeID int64,
) (*view.SalaryHistoryResponse, error) {

	if err := utility.ValidateEmployeeID(
		employeeID,
	); err != nil {

		return nil, err
	}

	history, err :=
		s.salaryRepository.GetSalaryHistory(
			ctx,
			employeeID,
		)

	if err != nil {
		return nil, err
	}

	return &view.SalaryHistoryResponse{
		EmployeeID: employeeID,
		History:    history,
	}, nil
}
