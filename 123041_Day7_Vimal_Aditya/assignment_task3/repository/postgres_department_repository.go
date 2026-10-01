package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"example.com/employee-management/model"
)

var ErrDepartmentNotFound = errors.New("department not found")

type PostgresDepartmentRepository struct {
	db *pgxpool.Pool
}

func NewPostgresDepartmentRepository(
	db *pgxpool.Pool,
) *PostgresDepartmentRepository {

	return &PostgresDepartmentRepository{
		db: db,
	}
}

func (r *PostgresDepartmentRepository) Save(
	department *model.Department,
) error {

	fmt.Println("\n----- CREATE DEPARTMENT -----")

	query := `
		INSERT INTO departments
		(name, code)
		VALUES ($1, $2)
		RETURNING id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		department.Name,
		department.Code,
	).Scan(&department.ID)

	if err != nil {
		return err
	}

	fmt.Println("Department created successfully")
	fmt.Println("Generated ID:", department.ID)

	return nil
}

func (r *PostgresDepartmentRepository) FindByID(id int64) (model.Department, error) {

	fmt.Println("\n----- READ DEPARTMENT -----")

	var department model.Department

	query := `
		SELECT
			id,
			name,
			code
		FROM departments
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&department.ID,
		&department.Name,
		&department.Code,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return department, ErrDepartmentNotFound
		}

		return department, err
	}

	fmt.Println("ID   :", department.ID)
	fmt.Println("Name :", department.Name)
	fmt.Println("Code :", department.Code)

	return department, nil
}

func (r *PostgresDepartmentRepository) FindAll() (
	[]model.Department,
	error,
) {

	fmt.Println("\n----- READ ALL DEPARTMENTS -----")

	query := `
		SELECT
			id,
			name,
			code
		FROM departments
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var departments []model.Department

	for rows.Next() {

		var department model.Department

		err := rows.Scan(
			&department.ID,
			&department.Name,
			&department.Code,
		)

		if err != nil {
			return nil, err
		}

		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *PostgresDepartmentRepository) Update(department model.Department) error {

	fmt.Println("\n----- UPDATE DEPARTMENT -----")

	query := `
		UPDATE departments
		SET
			name = $1,
			code = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		department.Name,
		department.Code,
		department.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrDepartmentNotFound
	}

	fmt.Println("Department updated successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *PostgresDepartmentRepository) Delete(id int64) error {

	fmt.Println("\n----- DELETE DEPARTMENT -----")

	result, err := r.db.Exec(
		context.Background(),
		`DELETE FROM departments WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrDepartmentNotFound
	}

	fmt.Println("Department deleted successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}