package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"employee-api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	List(ctx context.Context, department, sortBy, order string, limit, offset int) ([]model.Employee, error)
	Get(ctx context.Context, id int64) (model.Employee, error)
	Create(ctx context.Context, e model.Employee) (model.Employee, error)
	Update(ctx context.Context, id int64, e model.Employee) (model.Employee, error)
	Delete(ctx context.Context, id int64) error
}

type PostgresEmployeeRepository struct {
	DB *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &PostgresEmployeeRepository{DB: db}
}

func (r *PostgresEmployeeRepository) List(ctx context.Context, department, sortBy, order string, limit, offset int) ([]model.Employee, error) {
	columns := map[string]string{
		"id": "id", "name": "name", "salary": "salary", "department": "department",
	}
	column, ok := columns[sortBy]
	if !ok {
		column = "id"
	}
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	// Column and order are selected from allowlists above; user input is not inserted directly.
	query := fmt.Sprintf(`
        SELECT id, name, email, department, salary
        FROM employees
        WHERE ($1 = '' OR department ILIKE $1)
        ORDER BY %s %s
        LIMIT $2 OFFSET $3`, column, order)

	rows, err := r.DB.Query(ctx, query, "%"+department+"%", limit, offset)
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

func (r *PostgresEmployeeRepository) Get(ctx context.Context, id int64) (model.Employee, error) {
	var e model.Employee
	err := r.DB.QueryRow(ctx, `
        SELECT id, name, email, department, salary
        FROM employees WHERE id = $1`, id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, ErrEmployeeNotFound
	}
	return e, err
}

func (r *PostgresEmployeeRepository) Create(ctx context.Context, e model.Employee) (model.Employee, error) {
	err := r.DB.QueryRow(ctx, `
        INSERT INTO employees (name, email, department, salary)
        VALUES ($1, $2, $3, $4)
        RETURNING id, name, email, department, salary`,
		strings.TrimSpace(e.Name), strings.TrimSpace(e.Email),
		strings.TrimSpace(e.Department), e.Salary,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	return e, err
}

func (r *PostgresEmployeeRepository) Update(ctx context.Context, id int64, e model.Employee) (model.Employee, error) {
	err := r.DB.QueryRow(ctx, `
        UPDATE employees
        SET name=$1, email=$2, department=$3, salary=$4
        WHERE id=$5
        RETURNING id, name, email, department, salary`,
		strings.TrimSpace(e.Name), strings.TrimSpace(e.Email),
		strings.TrimSpace(e.Department), e.Salary, id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, ErrEmployeeNotFound
	}
	return e, err
}

func (r *PostgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.DB.Exec(ctx, `DELETE FROM employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}
	return nil
}
