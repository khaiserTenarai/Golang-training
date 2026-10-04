package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management-app/model"
)

var ErrDepartmentNotFound = errors.New("department not found")

type DepartmentRepositoryImpl struct {
	db *pgxpool.Pool
}

// Constructor
func NewDepartmentRepository(
	db *pgxpool.Pool,
) DepartmentRepository {

	return &DepartmentRepositoryImpl{
		db: db,
	}
}

// ==================================================
// CREATE
// ==================================================

func (r *DepartmentRepositoryImpl) Save(
	department model.Department,
) error {

	fmt.Println("\n----- CREATE DEPARTMENT -----")

	query := `
		INSERT INTO departments
			(name)
		VALUES
			($1)
		RETURNING id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		department.Name,
	).Scan(&department.ID)

	if err != nil {
		return err
	}

	fmt.Println("Department created successfully.")
	fmt.Println("Generated ID:", department.ID)

	return nil
}

// ==================================================
// READ ONE
// ==================================================

func (r *DepartmentRepositoryImpl) FindByID(
	id int,
) (model.Department, error) {

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
			return department, ErrDepartmentNotFound
		}

		return department, err
	}

	return department, nil
}

// ==================================================
// READ ALL
// ==================================================

func (r *DepartmentRepositoryImpl) FindAll() []model.Department {

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
		fmt.Println("Error fetching departments:", err)
		return []model.Department{}
	}

	defer rows.Close()

	departments := make([]model.Department, 0)

	for rows.Next() {

		var department model.Department

		err := rows.Scan(
			&department.ID,
			&department.Name,
		)

		if err != nil {
			fmt.Println("Error scanning department:", err)
			return []model.Department{}
		}

		departments = append(
			departments,
			department,
		)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error reading departments:", err)
		return []model.Department{}
	}

	return departments
}

// ==================================================
// UPDATE
// ==================================================

func (r *DepartmentRepositoryImpl) Update(
	department model.Department,
) error {

	fmt.Println("\n----- UPDATE DEPARTMENT -----")

	query := `
		UPDATE departments
		SET
			name = $1
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

	fmt.Println("Department updated successfully.")
	fmt.Println("Department ID:", department.ID)

	return nil
}

// ==================================================
// DELETE
// ==================================================

func (r *DepartmentRepositoryImpl) Delete(
	id int,
) error {

	fmt.Println("\n----- DELETE DEPARTMENT -----")

	query := `
		DELETE FROM departments
		WHERE id = $1
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrDepartmentNotFound
	}

	fmt.Println("Department deleted successfully.")
	fmt.Println("Department ID:", id)

	return nil
}
