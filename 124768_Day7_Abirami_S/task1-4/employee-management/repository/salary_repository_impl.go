package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrSalaryEmployeeNotFound = errors.New("employee not found")

type SalaryRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewSalaryRepository(db *pgxpool.Pool) SalaryRepository {
	return &SalaryRepositoryImpl{
		db: db,
	}
}

func (r *SalaryRepositoryImpl) UpdateSalary(employeeID int, newSalary float64) error {
	fmt.Println("\n----- UPDATE SALARY -----")

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	var oldSalary float64

	err = tx.QueryRow(
		context.Background(),
		`SELECT salary FROM employees WHERE id = $1`,
		employeeID,
	).Scan(&oldSalary)

	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrSalaryEmployeeNotFound
		}
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		`UPDATE employees SET salary = $1 WHERE id = $2`,
		newSalary,
		employeeID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		`INSERT INTO salary_history (employee_id, old_salary, new_salary)
		 VALUES ($1, $2, $3)`,
		employeeID,
		oldSalary,
		newSalary,
	)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Salary updated successfully")
	fmt.Println("Old Salary:", oldSalary)
	fmt.Println("New Salary:", newSalary)

	return nil
}
