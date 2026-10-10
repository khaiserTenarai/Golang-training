// Package repository implements PostgreSQL persistence.
package repository

import (
	"context"
	"employee-management/model"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

// EmployeeRepository defines persistence operations.
type EmployeeRepository interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string) ([]model.Employee, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
}

// PostgresEmployeeRepository is the PostgreSQL implementation.
type PostgresEmployeeRepository struct{ db *pgxpool.Pool }

// NewPostgresEmployeeRepository creates the repository.
func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &PostgresEmployeeRepository{db}
}

// Create inserts an employee.
func (r *PostgresEmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	return r.db.QueryRow(ctx, `INSERT INTO employees(name,email,department,salary) VALUES($1,$2,$3,$4) RETURNING id`, e.Name, e.Email, e.Department, e.Salary).Scan(&e.ID)
}

// GetAll returns all employees. Basic project supports only sort=name.
func (r *PostgresEmployeeRepository) GetAll(ctx context.Context, sortBy string) ([]model.Employee, error) {
	order := "id"
	if sortBy == "name" {
		order = "name"
	}
	rows, err := r.db.Query(ctx, "SELECT id,name,email,department,salary FROM employees ORDER BY "+order)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []model.Employee{}
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// GetByID retrieves one employee.
func (r *PostgresEmployeeRepository) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, "SELECT id,name,email,department,salary FROM employees WHERE id=$1", id).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// Update changes an employee.
func (r *PostgresEmployeeRepository) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	var u model.Employee
	err := r.db.QueryRow(ctx, `UPDATE employees SET name=$1,email=$2,department=$3,salary=$4 WHERE id=$5 RETURNING id,name,email,department,salary`, e.Name, e.Email, e.Department, e.Salary, id).Scan(&u.ID, &u.Name, &u.Email, &u.Department, &u.Salary)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Delete removes an employee.
func (r *PostgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.Exec(ctx, "DELETE FROM employees WHERE id=$1", id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("employee not found")
	}
	return nil
}

// IsNotFound identifies pgx no-row errors.
func IsNotFound(err error) bool { return errors.Is(err, pgx.ErrNoRows) }

// SafeSort is kept available for future repository extensions.
func SafeSort(v string) string {
	if strings.TrimSpace(v) == "" {
		return "id"
	}
	if v == "name" {
		return "name"
	}
	return "id"
}
