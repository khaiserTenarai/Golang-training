package repository

import (
	"context"
	"errors"
	"fmt"

	"employee-api-day13/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) *PostgresEmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}

// Task 11: every database operation accepts and forwards context.Context.
func (r *PostgresEmployeeRepository) Create(ctx context.Context, input model.EmployeeInput) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `
		INSERT INTO employees (name, email, department)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, department, created_at`,
		input.Name, input.Email, input.Department,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.CreatedAt)
	if err != nil {
		return model.Employee{}, fmt.Errorf("create employee: %w", err)
	}
	return e, nil
}

func (r *PostgresEmployeeRepository) List(ctx context.Context) ([]model.Employee, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email, department, created_at FROM employees ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("list employees: %w", err)
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan employee: %w", err)
		}
		employees = append(employees, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate employees: %w", err)
	}
	return employees, nil
}

func (r *PostgresEmployeeRepository) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `
		SELECT id, name, email, department, created_at FROM employees WHERE id=$1`, id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, ErrNotFound
	}
	if err != nil {
		return model.Employee{}, fmt.Errorf("get employee: %w", err)
	}
	return e, nil
}

func (r *PostgresEmployeeRepository) Update(ctx context.Context, id int64, input model.EmployeeInput) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `
		UPDATE employees SET name=$1, email=$2, department=$3
		WHERE id=$4
		RETURNING id, name, email, department, created_at`,
		input.Name, input.Email, input.Department, id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, ErrNotFound
	}
	if err != nil {
		return model.Employee{}, fmt.Errorf("update employee: %w", err)
	}
	return e, nil
}

func (r *PostgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM employees WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

var ErrNotFound = errors.New("employee not found")
