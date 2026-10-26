package repository

import (
	"context"
	"employee-service/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	Create(ctx context.Context, emp *entity.Employee) error
	FindByID(ctx context.Context, id int64) (*entity.Employee, error)
	FindAll(ctx context.Context) ([]entity.Employee, error)
	Update(ctx context.Context, emp *entity.Employee) error
	Delete(ctx context.Context, id int64) error
}

type pgxEmployeeRepo struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) EmployeeRepository {
	return &pgxEmployeeRepo{pool: pool}
}

func (r *pgxEmployeeRepo) Create(ctx context.Context, emp *entity.Employee) error {
	query := `INSERT INTO employees (first_name, last_name, email, department_id, status)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	return r.pool.QueryRow(ctx, query, emp.FirstName, emp.LastName, emp.Email, emp.DepartmentID, emp.Status).
		Scan(&emp.ID, &emp.CreatedAt, &emp.UpdatedAt)
}

func (r *pgxEmployeeRepo) FindByID(ctx context.Context, id int64) (*entity.Employee, error) {
	query := `SELECT id, first_name, last_name, email, department_id, status, created_at, updated_at FROM employees WHERE id = $1`
	var emp entity.Employee
	err := r.pool.QueryRow(ctx, query, id).Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Email, &emp.DepartmentID, &emp.Status, &emp.CreatedAt, &emp.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *pgxEmployeeRepo) FindAll(ctx context.Context) ([]entity.Employee, error) {
	query := `SELECT id, first_name, last_name, email, department_id, status, created_at, updated_at FROM employees`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []entity.Employee
	for rows.Next() {
		var emp entity.Employee
		if err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Email, &emp.DepartmentID, &emp.Status, &emp.CreatedAt, &emp.UpdatedAt); err == nil {
			emps = append(emps, emp)
		}
	}
	return emps, nil
}

func (r *pgxEmployeeRepo) Update(ctx context.Context, emp *entity.Employee) error {
	query := `UPDATE employees SET first_name=$1, last_name=$2, email=$3, department_id=$4, status=$5, updated_at=NOW() WHERE id=$6`
	_, err := r.pool.Exec(ctx, query, emp.FirstName, emp.LastName, emp.Email, emp.DepartmentID, emp.Status, emp.ID)
	return err
}

func (r *pgxEmployeeRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
