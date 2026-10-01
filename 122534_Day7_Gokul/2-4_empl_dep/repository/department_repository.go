// Package repository has all the direct database access. Every SQL
// query in this project lives here, so the rest of the code never has
// to deal with database/sql directly.
package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	"employee-department-service/model"
)

// ErrDepartmentNotFound is returned when a department ID doesn't exist.
var ErrDepartmentNotFound = errors.New("department not found")

// ErrDepartmentHasEmployees is returned when trying to delete a
// department that still has employees pointing to it. Postgres itself
// blocks the delete (that's what "ON DELETE RESTRICT" in schema.sql
// does) - we just turn its low-level error into something readable.
var ErrDepartmentHasEmployees = errors.New("department has employees and cannot be deleted")

// foreignKeyViolation is the Postgres error code for a foreign key
// constraint failure. See: https://www.postgresql.org/docs/current/errcodes-appendix.html
const foreignKeyViolation = "23503"

// CreateDepartment inserts a new department and returns its generated ID.
func CreateDepartment(db *sql.DB, name string) (int, error) {
	var id int
	query := `INSERT INTO department (name) VALUES ($1) RETURNING id`
	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create department failed: %w", err)
	}
	return id, nil
}

// GetDepartment fetches one department by ID.
func GetDepartment(db *sql.DB, id int) (model.Department, error) {
	var d model.Department
	query := `SELECT id, name, created_at FROM department WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&d.ID, &d.Name, &d.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Department{}, fmt.Errorf("get department failed: %w", ErrDepartmentNotFound)
		}
		return model.Department{}, fmt.Errorf("get department failed: %w", err)
	}
	return d, nil
}

// ListDepartments returns every department, ordered by name.
func ListDepartments(db *sql.DB) ([]model.Department, error) {
	query := `SELECT id, name, created_at FROM department ORDER BY name`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("list departments failed: %w", err)
	}
	defer rows.Close()

	var departments []model.Department
	for rows.Next() {
		var d model.Department
		if err := rows.Scan(&d.ID, &d.Name, &d.CreatedAt); err != nil {
			return nil, fmt.Errorf("list departments failed: %w", err)
		}
		departments = append(departments, d)
	}
	return departments, rows.Err()
}

// UpdateDepartment renames an existing department.
func UpdateDepartment(db *sql.DB, id int, newName string) error {
	query := `UPDATE department SET name = $1 WHERE id = $2`
	result, err := db.Exec(query, newName, id)
	if err != nil {
		return fmt.Errorf("update department failed: %w", err)
	}
	return checkRowsAffected(result, ErrDepartmentNotFound)
}

// DeleteDepartment removes a department by ID. If any employee still
// belongs to it, Postgres refuses the delete with a foreign key
// error - we detect that specific error and return the friendlier
// ErrDepartmentHasEmployees instead.
func DeleteDepartment(db *sql.DB, id int) error {
	query := `DELETE FROM department WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == foreignKeyViolation {
			return fmt.Errorf("delete department failed: %w", ErrDepartmentHasEmployees)
		}
		return fmt.Errorf("delete department failed: %w", err)
	}
	return checkRowsAffected(result, ErrDepartmentNotFound)
}

// checkRowsAffected is a small shared helper: if an UPDATE/DELETE
// touched zero rows, the ID we were given didn't exist.
func checkRowsAffected(result sql.Result, notFoundErr error) error {
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check affected rows: %w", err)
	}
	if rows == 0 {
		return notFoundErr
	}
	return nil
}
