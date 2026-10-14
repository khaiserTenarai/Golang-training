package repository

import (
	"context"
	"errors"

	"employee-rest-api/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("employee not found")
var ErrEmailExists = errors.New("an employee with this email already exists")

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	err := r.db.QueryRow(ctx, `
		INSERT INTO employees (name, email, department, salary)
		VALUES ($1, $2, $3, $4) RETURNING id
	`, e.Name, e.Email, e.Department, e.Salary).Scan(&e.ID)
	return mapDBError(err)
}

func (r *EmployeeRepository) GetAll(ctx context.Context) ([]model.Employee, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, email, department, salary
		FROM employees ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, rows.Err()
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `
		SELECT id, name, email, department, salary
		FROM employees WHERE id = $1
	`, id).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return e, ErrNotFound
	}
	return e, err
}

func (r *EmployeeRepository) Update(ctx context.Context, id int64, e *model.Employee) (model.Employee, error) {
	var updated model.Employee
	err := r.db.QueryRow(ctx, `
		UPDATE employees
		SET name = $1, email = $2, department = $3, salary = $4
		WHERE id = $5
		RETURNING id, name, email, department, salary
	`, e.Name, e.Email, e.Department, e.Salary, id).Scan(
		&updated.ID, &updated.Name, &updated.Email, &updated.Department, &updated.Salary,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return updated, ErrNotFound
	}
	if err != nil {
		return updated, mapDBError(err)
	}
	return updated, nil
}

func (r *EmployeeRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func mapDBError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return ErrEmailExists
	}
	return err
}
