package repository

import (
	"context"
	"employee-management/model"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrDepartmentNotFound = errors.New("department not found")

type DepartmentRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewDepartmentRepository(db *pgxpool.Pool) DepartmentRepository {
	return &DepartmentRepositoryImpl{
		db: db,
	}
}
func (r *DepartmentRepositoryImpl) CreateDepartment(department model.Department) error {
	fmt.Println("\n----- CREATE DEPARTMENT -----")

	query := `
		INSERT INTO departments
		(name)
		VALUES ($1)
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		department.Name,
	)

	if err != nil {
		return err
	}

	fmt.Println("Department created successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}
func (r *DepartmentRepositoryImpl) GetDepartment(id int) (*model.Department, error) {
	fmt.Println("\n----- READ DEPARTMENT -----")

	var department model.Department

	query := `
		SELECT
			id,
			name
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
	)

	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, ErrDepartmentNotFound
		}

		return nil, err
	}

	return &department, nil
}
func (r *DepartmentRepositoryImpl) GetAllDepartments() ([]model.Department, error) {
	fmt.Println("\n----- READ ALL DEPARTMENTS -----")

	query := `
		SELECT
			id,
			name
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
func (r *DepartmentRepositoryImpl) UpdateDepartment(department model.Department) error {
	fmt.Println("\n----- UPDATE DEPARTMENT -----")

	query := `
		UPDATE departments
		SET name = $1
		WHERE id = $2
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		department.Name,
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
func (r *DepartmentRepositoryImpl) DeleteDepartment(id int) error {
	fmt.Println("\n----- DELETE DEPARTMENT -----")

	result, err := r.db.Exec(
		context.Background(),
		"DELETE FROM departments WHERE id = $1",
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
