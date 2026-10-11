// Package repository manages database operations.
package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/model"
)

var ErrNotFound = errors.New("employee not found")

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

type EmployeeRepository interface {
	Create(ctx context.Context, emp *model.Employee) error
	GetAll(ctx context.Context, sort string) ([]model.Employee, error)
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	Update(ctx context.Context, id int64, emp *model.Employee) (*model.Employee, error)
	PartialUpdate(ctx context.Context, id int64, patch *model.PartialEmployee) (*model.Employee, error)
	Delete(ctx context.Context, id int64) error
}

type postgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &postgresEmployeeRepository{db: db}
}

func (r *postgresEmployeeRepository) Create(ctx context.Context, emp *model.Employee) error {
	query := `INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(ctx, query, emp.Name, emp.Email, emp.Department, emp.Salary).Scan(&emp.ID)
}

func (r *postgresEmployeeRepository) GetAll(ctx context.Context, sort string) ([]model.Employee, error) {
	sortBy := "id"
	order := "ASC"

	if strings.HasPrefix(sort, "-") {
		order = "DESC"
		sort = strings.TrimPrefix(sort, "-")
	}

	if sort == "name" || sort == "salary" || sort == "department" {
		sortBy = sort
	}

	query := fmt.Sprintf("SELECT id, name, email, department, salary FROM employees ORDER BY %s %s", sortBy, order)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func (r *postgresEmployeeRepository) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	query := `SELECT id, name, email, department, salary FROM employees WHERE id = $1`
	var e model.Employee
	err := r.db.QueryRow(ctx, query, id).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *postgresEmployeeRepository) Update(ctx context.Context, id int64, emp *model.Employee) (*model.Employee, error) {
	query := `UPDATE employees SET name = $1, email = $2, department = $3, salary = $4 WHERE id = $5 RETURNING id, name, email, department, salary`
	var updated model.Employee
	err := r.db.QueryRow(ctx, query, emp.Name, emp.Email, emp.Department, emp.Salary, id).Scan(&updated.ID, &updated.Name, &updated.Email, &updated.Department, &updated.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *postgresEmployeeRepository) PartialUpdate(ctx context.Context, id int64, patch *model.PartialEmployee) (*model.Employee, error) {
	current, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if patch.Name != nil {
		current.Name = *patch.Name
	}
	if patch.Email != nil {
		current.Email = *patch.Email
	}
	if patch.Department != nil {
		current.Department = *patch.Department
	}
	if patch.Salary != nil {
		current.Salary = *patch.Salary
	}

	return r.Update(ctx, id, current)
}

func (r *postgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM employees WHERE id = $1`
	cmd, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}