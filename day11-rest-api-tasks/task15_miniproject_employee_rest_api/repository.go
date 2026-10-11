package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

const schema = `
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    department VARCHAR(100) NOT NULL,
    salary NUMERIC(12, 2) NOT NULL CHECK (salary > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_employees_department ON employees (LOWER(department));
`

const columns = "id, name, email, department, salary, created_at, updated_at"

var sortColumns = map[string]string{
	"id":         "id",
	"name":       "name",
	"salary":     "salary",
	"department": "department",
	"created_at": "created_at",
}

var likeEscaper = strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func Migrate(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, schema)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanEmployee(s scanner) (Employee, error) {
	var e Employee
	err := s.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func mapError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return ErrDuplicateEmail
	}
	return err
}

func (r *PostgresRepository) Create(ctx context.Context, in EmployeeInput) (Employee, error) {
	row := r.db.QueryRowContext(ctx,
		`INSERT INTO employees (name, email, department, salary)
		 VALUES ($1, $2, $3, $4)
		 RETURNING `+columns,
		in.Name, in.Email, in.Department, in.Salary)

	e, err := scanEmployee(row)
	if err != nil {
		return Employee{}, mapError(err)
	}
	return e, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int) (Employee, error) {
	row := r.db.QueryRowContext(ctx, `SELECT `+columns+` FROM employees WHERE id = $1`, id)

	e, err := scanEmployee(row)
	if err != nil {
		return Employee{}, mapError(err)
	}
	return e, nil
}

func (r *PostgresRepository) List(ctx context.Context, p ListParams) ([]Employee, int, error) {
	var conditions []string
	var args []any

	if p.Department != "" {
		args = append(args, p.Department)
		conditions = append(conditions, fmt.Sprintf("LOWER(department) = LOWER($%d)", len(args)))
	}
	if p.Search != "" {
		args = append(args, "%"+likeEscaper.Replace(p.Search)+"%")
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR email ILIKE $%d)", len(args), len(args)))
	}

	where := ""
	if len(conditions) > 0 {
		where = " WHERE " + strings.Join(conditions, " AND ")
	}

	var total int
	if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM employees"+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortColumn, ok := sortColumns[p.Sort]
	if !ok {
		sortColumn = "id"
	}
	direction := "ASC"
	if strings.EqualFold(p.Order, "desc") {
		direction = "DESC"
	}

	args = append(args, p.Limit, (p.Page-1)*p.Limit)
	query := fmt.Sprintf(
		"SELECT %s FROM employees%s ORDER BY %s %s, id ASC LIMIT $%d OFFSET $%d",
		columns, where, sortColumn, direction, len(args)-1, len(args))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		e, err := scanEmployee(rows)
		if err != nil {
			return nil, 0, err
		}
		employees = append(employees, e)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

func (r *PostgresRepository) Update(ctx context.Context, id int, in EmployeeInput) (Employee, error) {
	row := r.db.QueryRowContext(ctx,
		`UPDATE employees
		 SET name = $1, email = $2, department = $3, salary = $4, updated_at = NOW()
		 WHERE id = $5
		 RETURNING `+columns,
		in.Name, in.Email, in.Department, in.Salary, id)

	e, err := scanEmployee(row)
	if err != nil {
		return Employee{}, mapError(err)
	}
	return e, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *PostgresRepository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}
