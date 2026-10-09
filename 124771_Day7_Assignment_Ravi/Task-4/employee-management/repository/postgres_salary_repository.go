package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/model"
)

var ErrEmployeeNotFound = errors.New(
	"employee not found",
)

type PostgresSalaryRepository struct {
	db *pgxpool.Pool
}

func NewPostgresSalaryRepository(
	db *pgxpool.Pool,
) *PostgresSalaryRepository {

	return &PostgresSalaryRepository{
		db: db,
	}
}

// ----------------------------------------------------
// Get Employee By ID
// ----------------------------------------------------

func (r *PostgresSalaryRepository) GetByID(
	ctx context.Context,
	tx pgx.Tx,
	employeeID int64,
) (*model.Employee, error) {

	var employee model.Employee

	err := tx.QueryRow(
		ctx,
		`
		SELECT
			id,
			name,
			department,
			salary
		FROM employees
		WHERE id = $1
		FOR UPDATE
		`,
		employeeID,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Department,
		&employee.Salary,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEmployeeNotFound
		}

		return nil, err
	}

	return &employee, nil
}

// ----------------------------------------------------
// Update Employee Salary
// ----------------------------------------------------

func (r *PostgresSalaryRepository) UpdateSalary(
	ctx context.Context,
	tx pgx.Tx,
	employeeID int64,
	newSalary float64,
) error {

	_, err := tx.Exec(
		ctx,
		`
		UPDATE employees
		SET salary = $1
		WHERE id = $2
		`,
		newSalary,
		employeeID,
	)

	return err
}

// ----------------------------------------------------
// Create Salary History
// ----------------------------------------------------

func (r *PostgresSalaryRepository) CreateSalaryHistory(
	ctx context.Context,
	tx pgx.Tx,
	employeeID int64,
	oldSalary float64,
	newSalary float64,
) (*model.SalaryHistory, error) {

	var history model.SalaryHistory

	err := tx.QueryRow(
		ctx,
		`
		INSERT INTO salary_history (
			employee_id,
			old_salary,
			new_salary
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			employee_id,
			old_salary,
			new_salary,
			changed_at
		`,
		employeeID,
		oldSalary,
		newSalary,
	).Scan(
		&history.ID,
		&history.EmployeeID,
		&history.OldSalary,
		&history.NewSalary,
		&history.ChangedAt,
	)

	if err != nil {
		return nil, err
	}

	return &history, nil
}

// ----------------------------------------------------
// Get Salary History
// ----------------------------------------------------

func (r *PostgresSalaryRepository) GetSalaryHistory(
	ctx context.Context,
	employeeID int64,
) ([]model.SalaryHistory, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			employee_id,
			old_salary,
			new_salary,
			changed_at
		FROM salary_history
		WHERE employee_id = $1
		ORDER BY changed_at DESC, id DESC
		`,
		employeeID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	historyList :=
		make([]model.SalaryHistory, 0)

	for rows.Next() {

		var history model.SalaryHistory

		err := rows.Scan(
			&history.ID,
			&history.EmployeeID,
			&history.OldSalary,
			&history.NewSalary,
			&history.ChangedAt,
		)

		if err != nil {
			return nil, err
		}

		historyList = append(
			historyList,
			history,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return historyList, nil
}
