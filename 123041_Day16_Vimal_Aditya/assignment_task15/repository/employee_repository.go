package repository

import (
	"context"
	"errors"

	"employee-management/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	Create(ctx context.Context, e *model.Employee) error
	GetAll(ctx context.Context, sortBy string) ([]model.Employee, error)
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error)
	Delete(ctx context.Context, id int64) error
}

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}

func (r *PostgresEmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO employees (name, email, department_id, department_name, salary) 
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		e.Name, e.Email, e.Department.ID, e.Department.Name, e.Salary).Scan(&e.ID)
}

func (r *PostgresEmployeeRepository) GetAll(ctx context.Context, sortBy string) ([]model.Employee, error) {
	orderBy := "id"
	if sortBy == "name" {
		orderBy = "name"
	}
	rows, err := r.db.Query(ctx, "SELECT id, name, email, department_id, department_name, salary FROM employees ORDER BY "+orderBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department.ID, &e.Department.Name, &e.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, rows.Err()
}

func (r *PostgresEmployeeRepository) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, 
		"SELECT id, name, email, department_id, department_name, salary FROM employees WHERE id = $1", id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department.ID, &e.Department.Name, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("employee not found")
	}
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *PostgresEmployeeRepository) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	var updated model.Employee
	err := r.db.QueryRow(ctx,
		`UPDATE employees SET name=$1, email=$2, department_id=$3, department_name=$4, salary=$5 
		 WHERE id=$6 RETURNING id, name, email, department_id, department_name, salary`,
		e.Name, e.Email, e.Department.ID, e.Department.Name, e.Salary, id).
		Scan(&updated.ID, &updated.Name, &updated.Email, &updated.Department.ID, &updated.Department.Name, &updated.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("employee not found")
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PostgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.Exec(ctx, "DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("employee not found")
	}
	return nil
}