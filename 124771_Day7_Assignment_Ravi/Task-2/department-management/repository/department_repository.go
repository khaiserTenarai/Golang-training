package repository

import (
	"context"
	"fmt"

	"department-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository struct {
	DB *pgxpool.Pool
}

func NewDepartmentRepository(db *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{
		DB: db,
	}
}

func (r *DepartmentRepository) Create(
	ctx context.Context,
	department *model.Department,
) error {

	query := `
		INSERT INTO departments (name, description)
		VALUES ($1, $2)
		RETURNING id
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		department.Name,
		department.Description,
	).Scan(&department.ID)

	if err != nil {
		return fmt.Errorf("failed to create department: %w", err)
	}

	return nil
}

func (r *DepartmentRepository) GetByID(
	ctx context.Context,
	id int,
) (*model.Department, error) {

	query := `
		SELECT id, name, description
		FROM departments
		WHERE id = $1
	`

	var department model.Department

	err := r.DB.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&department.ID,
		&department.Name,
		&department.Description,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("department not found")
		}

		return nil, fmt.Errorf(
			"failed to get department: %w",
			err,
		)
	}

	return &department, nil
}

func (r *DepartmentRepository) GetAll(
	ctx context.Context,
) ([]model.Department, error) {

	query := `
		SELECT id, name, description
		FROM departments
		ORDER BY id
	`

	rows, err := r.DB.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to get departments: %w",
			err,
		)
	}

	defer rows.Close()

	departments := make([]model.Department, 0)

	for rows.Next() {

		var department model.Department

		err := rows.Scan(
			&department.ID,
			&department.Name,
			&department.Description,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to scan department: %w",
				err,
			)
		}

		departments = append(
			departments,
			department,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"failed while reading departments: %w",
			err,
		)
	}

	return departments, nil
}

func (r *DepartmentRepository) Update(
	ctx context.Context,
	department *model.Department,
) error {

	query := `
		UPDATE departments
		SET name = $1,
		    description = $2
		WHERE id = $3
	`

	result, err := r.DB.Exec(
		ctx,
		query,
		department.Name,
		department.Description,
		department.ID,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to update department: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("department not found")
	}

	return nil
}

func (r *DepartmentRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
		DELETE FROM departments
		WHERE id = $1
	`

	result, err := r.DB.Exec(
		ctx,
		query,
		id,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to delete department: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("department not found")
	}

	return nil
}
