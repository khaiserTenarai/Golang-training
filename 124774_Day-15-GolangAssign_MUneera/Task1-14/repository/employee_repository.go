package repository

import (
	"context"
	"errors"
	"fmt"

	"employee-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, int, int, string, string) ([]model.Employee, int, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
}

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(
	db *pgxpool.Pool,
) EmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}

func (r *PostgresEmployeeRepository) Create(
	ctx context.Context,
	e *model.Employee,
) error {
	query := `
        INSERT INTO employees (name, email, department, salary)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	return r.db.QueryRow(
		ctx,
		query,
		e.Name,
		e.Email,
		e.Department,
		e.Salary,
	).Scan(&e.ID)
}

func (r *PostgresEmployeeRepository) GetAll(
	ctx context.Context,
	page int,
	size int,
	department string,
	sortBy string,
) ([]model.Employee, int, error) {
	where := ""
	args := []any{}

	if department != "" {
		where = " WHERE department = $1"
		args = append(args, department)
	}

	countQuery := "SELECT COUNT(*) FROM employees" + where

	var total int

	if err := r.db.QueryRow(
		ctx,
		countQuery,
		args...,
	).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Only allow known column names in the ORDER BY clause.
	orderBy := "id"

	if sortBy == "name" {
		orderBy = "name"
	}

	offset := (page - 1) * size

	query := fmt.Sprintf(`
        SELECT id, name, email, department, salary
        FROM employees
        %s
        ORDER BY %s
        LIMIT $%d OFFSET $%d
    `, where, orderBy, len(args)+1, len(args)+2)

	args = append(args, size, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	employees := []model.Employee{}

	for rows.Next() {
		var e model.Employee

		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Email,
			&e.Department,
			&e.Salary,
		); err != nil {
			return nil, 0, err
		}

		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

func (r *PostgresEmployeeRepository) GetByID(
	ctx context.Context,
	id int64,
) (*model.Employee, error) {
	var e model.Employee

	query := `
        SELECT id, name, email, department, salary
        FROM employees
        WHERE id = $1
    `

	err := r.db.QueryRow(ctx, query, id).Scan(
		&e.ID,
		&e.Name,
		&e.Email,
		&e.Department,
		&e.Salary,
	)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *PostgresEmployeeRepository) Update(
	ctx context.Context,
	id int64,
	e *model.Employee,
) (*model.Employee, error) {
	var updated model.Employee

	query := `
        UPDATE employees
        SET name = $1, email = $2, department = $3, salary = $4
        WHERE id = $5
        RETURNING id, name, email, department, salary
    `

	err := r.db.QueryRow(
		ctx,
		query,
		e.Name,
		e.Email,
		e.Department,
		e.Salary,
		id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Email,
		&updated.Department,
		&updated.Salary,
	)

	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *PostgresEmployeeRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	result, err := r.db.Exec(
		ctx,
		"DELETE FROM employees WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("employee not found")
	}

	return nil
}

func IsNotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
