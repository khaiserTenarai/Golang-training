package repository

import (
	"context"
	"employee-management/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	Save(ctx context.Context, emp *model.Employee) error
	FindAll(ctx context.Context, offset, limit int) ([]model.Employee, int, error)
	FindByID(ctx context.Context, id int64) (*model.Employee, error)
	UpdateAvatar(ctx context.Context, id int64, path string) error
	Delete(ctx context.Context, id int64) error
}

type postgresEmployeeRepo struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &postgresEmployeeRepo{db: db}
}

func (r *postgresEmployeeRepo) Save(ctx context.Context, emp *model.Employee) error {
	query := `INSERT INTO employees (name, email, department, designation, salary, avatar_path) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRow(ctx, query, emp.Name, emp.Email, emp.Department, emp.Designation, emp.Salary, emp.AvatarPath).Scan(&emp.ID)
}

func (r *postgresEmployeeRepo) FindAll(ctx context.Context, offset, limit int) ([]model.Employee, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM employees`).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, `SELECT id, name, email, department, designation, salary, COALESCE(avatar_path, '') FROM employees OFFSET $1 LIMIT $2`, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Designation, &emp.Salary, &emp.AvatarPath); err != nil {
			return nil, 0, err
		}
		employees = append(employees, emp)
	}
	return employees, total, nil
}

func (r *postgresEmployeeRepo) FindByID(ctx context.Context, id int64) (*model.Employee, error) {
	row := r.db.QueryRow(ctx, `SELECT id, name, email, department, designation, salary, COALESCE(avatar_path, '') FROM employees WHERE id = $1`, id)
	var emp model.Employee
	if err := row.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Designation, &emp.Salary, &emp.AvatarPath); err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *postgresEmployeeRepo) UpdateAvatar(ctx context.Context, id int64, path string) error {
	_, err := r.db.Exec(ctx, `UPDATE employees SET avatar_path = $1 WHERE id = $2`, path, id)
	return err
}

func (r *postgresEmployeeRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM employees WHERE id = $1`, id)
	return err
}