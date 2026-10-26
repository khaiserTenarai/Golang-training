package repository

import (
	"context"
	"encoding/json"

	"day18/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	CreateWithOutbox(ctx context.Context, emp *model.Employee, eventType string, payload map[string]interface{}) error
	FindByID(ctx context.Context, id int64) (*model.Employee, error)
	UpdateStatus(ctx context.Context, id int64, status model.EmployeeStatus) error
}

type pgxEmployeeRepo struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) EmployeeRepository {
	return &pgxEmployeeRepo{pool: pool}
}

func (r *pgxEmployeeRepo) CreateWithOutbox(ctx context.Context, emp *model.Employee, eventType string, payload map[string]interface{}) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Insert employee and fetch auto-incremented integer ID
	queryEmp := `INSERT INTO employees (first_name, last_name, email, department_id, status)
                 VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRow(ctx, queryEmp, emp.FirstName, emp.LastName, emp.Email, emp.DepartmentID, emp.Status).Scan(&emp.ID)
	if err != nil {
		return err
	}

	payload["employee_id"] = emp.ID
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	queryOutbox := `INSERT INTO outbox_events (aggregate_type, aggregate_id, event_type, payload)
                    VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, queryOutbox, "EMPLOYEE", emp.ID, eventType, jsonPayload)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *pgxEmployeeRepo) FindByID(ctx context.Context, id int64) (*model.Employee, error) {
	query := `SELECT id, first_name, last_name, email, department_id, status, created_at, updated_at 
              FROM employees WHERE id = $1`

	var emp model.Employee
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&emp.ID, &emp.FirstName, &emp.LastName, &emp.Email,
		&emp.DepartmentID, &emp.Status, &emp.CreatedAt, &emp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *pgxEmployeeRepo) UpdateStatus(ctx context.Context, id int64, status model.EmployeeStatus) error {
	query := `UPDATE employees SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.pool.Exec(ctx, query, status, id)
	return err
}