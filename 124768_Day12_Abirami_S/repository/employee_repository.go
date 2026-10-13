package repository

import (
	"context"
	"fmt"
	"strings"

	"day12/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct{ db *pgxpool.Pool }

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, in model.EmployeeInput) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `INSERT INTO employees(name,email,department,salary)
		VALUES($1,$2,$3,$4) RETURNING id,name,email,department,salary,created_at,updated_at`,
		in.Name, in.Email, in.Department, in.Salary).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `SELECT id,name,email,department,salary,created_at,updated_at FROM employees WHERE id=$1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (r *EmployeeRepository) GetAll(ctx context.Context, page, limit int, department, sortBy, order string) ([]model.Employee, int, error) {
	// sortBy and order are allow-listed before they are inserted into SQL.
	allowedSort := map[string]string{"id": "id", "name": "name", "salary": "salary", "department": "department", "created_at": "created_at"}
	column := allowedSort[sortBy]
	if column == "" {
		column = "id"
	}
	direction := "ASC"
	if strings.EqualFold(order, "desc") {
		direction = "DESC"
	}

	var total int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM employees WHERE ($1 = '' OR department ILIKE $2)`,
		department, "%"+department+"%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf(`SELECT id,name,email,department,salary,created_at,updated_at
		FROM employees WHERE ($1 = '' OR department ILIKE $2)
		ORDER BY %s %s LIMIT $3 OFFSET $4`, column, direction)
	rows, err := r.db.Query(ctx, query, department, "%"+department+"%", limit, (page-1)*limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, e)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *EmployeeRepository) Update(ctx context.Context, id int64, in model.EmployeeInput) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `UPDATE employees SET name=$1,email=$2,department=$3,salary=$4,updated_at=NOW()
		WHERE id=$5 RETURNING id,name,email,department,salary,created_at,updated_at`,
		in.Name, in.Email, in.Department, in.Salary, id).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (r *EmployeeRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *EmployeeRepository) Ping(ctx context.Context) error { return r.db.Ping(ctx) }
