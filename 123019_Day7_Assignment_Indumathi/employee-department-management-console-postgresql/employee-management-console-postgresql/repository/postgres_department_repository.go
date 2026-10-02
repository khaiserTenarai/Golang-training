package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"example.com/employee-management/models"
)

type PostgresDepartmentRepository struct {
	db *pgxpool.Pool
}

func NewPostgresDepartmentRepository(db *pgxpool.Pool) *PostgresDepartmentRepository {
	return &PostgresDepartmentRepository{db: db}
}

func (r *PostgresDepartmentRepository) Create(dept *models.Department) error {
	query := `
		INSERT INTO departments (name, code)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(context.Background(), query, dept.Name, dept.Code).
		Scan(&dept.ID, &dept.CreatedAt, &dept.UpdatedAt)
}

func (r *PostgresDepartmentRepository) FindByID(id int) (*models.Department, error) {
	query := `SELECT id, name, code, created_at, updated_at FROM departments WHERE id = $1`

	dept := &models.Department{}
	err := r.db.QueryRow(context.Background(), query, id).
		Scan(&dept.ID, &dept.Name, &dept.Code, &dept.CreatedAt, &dept.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("department not found")
		}
		return nil, err
	}
	return dept, nil
}

func (r *PostgresDepartmentRepository) FindAll() ([]models.Department, error) {
	query := `SELECT id, name, code, created_at, updated_at FROM departments ORDER BY id ASC`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []models.Department
	for rows.Next() {
		var d models.Department
		if err := rows.Scan(&d.ID, &d.Name, &d.Code, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		depts = append(depts, d)
	}
	return depts, nil
}

func (r *PostgresDepartmentRepository) Update(dept *models.Department) error {
	query := `
		UPDATE departments
		SET name = $1, code = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3`

	cmd, err := r.db.Exec(context.Background(), query, dept.Name, dept.Code, dept.ID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("department not found")
	}
	return nil
}

func (r *PostgresDepartmentRepository) Delete(id int) error {
	query := `DELETE FROM departments WHERE id = $1`

	cmd, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("department not found")
	}
	return nil
}