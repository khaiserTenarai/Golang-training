package repository

import (
	"context"
	"department-service/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository interface {
	Create(ctx context.Context, dept *entity.Department) error
	FindByID(ctx context.Context, id int64) (*entity.Department, error)
	FindAll(ctx context.Context) ([]entity.Department, error)
}

type pgxDepartmentRepo struct {
	pool *pgxpool.Pool
}

func NewDepartmentRepository(pool *pgxpool.Pool) DepartmentRepository {
	return &pgxDepartmentRepo{pool: pool}
}

func (r *pgxDepartmentRepo) Create(ctx context.Context, dept *entity.Department) error {
	query := `INSERT INTO departments (name, code, headcount_limit) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.pool.QueryRow(ctx, query, dept.Name, dept.Code, dept.HeadcountLimit).Scan(&dept.ID, &dept.CreatedAt)
}

func (r *pgxDepartmentRepo) FindByID(ctx context.Context, id int64) (*entity.Department, error) {
	query := `SELECT id, name, code, headcount_limit, current_headcount, created_at FROM departments WHERE id = $1`
	var dept entity.Department
	err := r.pool.QueryRow(ctx, query, id).Scan(&dept.ID, &dept.Name, &dept.Code, &dept.HeadcountLimit, &dept.CurrentHeadcount, &dept.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *pgxDepartmentRepo) FindAll(ctx context.Context) ([]entity.Department, error) {
	query := `SELECT id, name, code, headcount_limit, current_headcount, created_at FROM departments`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []entity.Department
	for rows.Next() {
		var dept entity.Department
		if err := rows.Scan(&dept.ID, &dept.Name, &dept.Code, &dept.HeadcountLimit, &dept.CurrentHeadcount, &dept.CreatedAt); err == nil {
			depts = append(depts, dept)
		}
	}
	return depts, nil
}
