package repository
import (
	"context"
	"errors"

	"salary-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SalaryRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewSalaryRepository(
	db *pgxpool.Pool,
) SalaryRepository {

	return &SalaryRepositoryImpl{
		db: db,
	}
}

func (r *SalaryRepositoryImpl) UpdateSalary(
	employeeID int,
	newSalary float64,
) error {

	ctx := context.Background()

	// Start transaction
	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	// Rollback if anything fails.
	defer tx.Rollback(ctx)

	var oldSalary float64

	// Get current salary.
	// FOR UPDATE locks this employee until
	// the transaction is committed.
	query := `
		SELECT salary
		FROM employees
		WHERE id = $1
		FOR UPDATE
	`

	err = tx.QueryRow(
		ctx,
		query,
		employeeID,
	).Scan(&oldSalary)

	if errors.Is(err, pgx.ErrNoRows) {
		return errors.New("employee not found")
	}

	if err != nil {
		return err
	}

	// Save salary history.
	query = `
		INSERT INTO salary_history
			(employee_id, old_salary, new_salary)
		VALUES
			($1, $2, $3)
	`

	_, err = tx.Exec(
		ctx,
		query,
		employeeID,
		oldSalary,
		newSalary,
	)

	if err != nil {
		return err
	}

	// Update employee salary.
	query = `
		UPDATE employees
		SET salary = $1
		WHERE id = $2
	`

	_, err = tx.Exec(
		ctx,
		query,
		newSalary,
		employeeID,
	)

	if err != nil {
		return err
	}

	// Commit transaction.
	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *SalaryRepositoryImpl) FindHistory(
	employeeID int,
) ([]model.SalaryHistory, error) {

	query := `
		SELECT
			id,
			employee_id,
			old_salary,
			new_salary,
			changed_at
		FROM salary_history
		WHERE employee_id = $1
		ORDER BY changed_at DESC
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		employeeID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	history := make(
		[]model.SalaryHistory,
		0,
	)

	for rows.Next() {

		var item model.SalaryHistory

		err := rows.Scan(
			&item.ID,
			&item.EmployeeID,
			&item.OldSalary,
			&item.NewSalary,
			&item.ChangedAt,
		)

		if err != nil {
			return nil, err
		}

		history = append(
			history,
			item,
		)
	}

	return history, rows.Err()
}
