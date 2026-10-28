package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"employee-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string, string, string, string, int, int) ([]model.Employee, int64, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
	Ping(context.Context) error
}

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}

func (r *PostgresEmployeeRepository) Ping(ctx context.Context) error {
	return r.db.Ping(ctx)
}

func (r *PostgresEmployeeRepository) Create(
	ctx context.Context,
	e *model.Employee,
) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO employees (name, email, department, salary)
         VALUES ($1, $2, $3, $4)
         RETURNING id`,
		e.Name, e.Email, e.Department, e.Salary,
	).Scan(&e.ID)
}

func (r *PostgresEmployeeRepository) GetAll(
	ctx context.Context,
	department, name, sortBy, sortOrder string,
	page, pageSize int,
) ([]model.Employee, int64, error) {
	// Only allow known SQL column names.
	allowedSort := map[string]string{
		"id":         "id",
		"name":       "name",
		"department": "department",
		"salary":     "salary",
	}

	orderBy, ok := allowedSort[strings.ToLower(sortBy)]
	if !ok {
		orderBy = "id"
	}

	direction := "ASC"
	if strings.EqualFold(sortOrder, "desc") {
		direction = "DESC"
	}

	var total int64

	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM employees
         WHERE ($1 = '' OR department ILIKE $1)
         AND ($2 = '' OR name ILIKE $2)`,
		department, name,
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	// orderBy and direction come only from the allowlists above.
	query := fmt.Sprintf(`
        SELECT id, name, email, department, salary
        FROM employees
        WHERE ($1 = '' OR department ILIKE $1)
          AND ($2 = '' OR name ILIKE $2)
        ORDER BY %s %s
        LIMIT $3 OFFSET $4`,
		orderBy, direction,
	)

	rows, err := r.db.Query(
		ctx, query, department, name, pageSize, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)

	for rows.Next() {
		var e model.Employee

		if err := rows.Scan(
			&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary,
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

	err := r.db.QueryRow(ctx,
		`SELECT id, name, email, department, salary
         FROM employees WHERE id = $1`,
		id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
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

	err := r.db.QueryRow(ctx,
		`UPDATE employees
         SET name=$1, email=$2, department=$3, salary=$4
         WHERE id=$5
         RETURNING id, name, email, department, salary`,
		e.Name, e.Email, e.Department, e.Salary, id,
	).Scan(
		&updated.ID, &updated.Name, &updated.Email,
		&updated.Department, &updated.Salary,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
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
		ctx, `DELETE FROM employees WHERE id=$1`, id,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
