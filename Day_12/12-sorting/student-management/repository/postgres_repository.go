package repository

import (
	"context"
	"fmt"
	"student-management/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStudentRepository struct{ DB *pgxpool.Pool }

func NewPostgresStudentRepository(db *pgxpool.Pool) StudentRepository {
	return &PostgresStudentRepository{DB: db}
}

func (r *PostgresStudentRepository) Create(ctx context.Context, student model.Student) error {
	_, err := r.DB.Exec(ctx, `INSERT INTO students (name, grade) VALUES ($1, $2)`, student.Name, student.Grade)
	return err
}

func (r *PostgresStudentRepository) GetFilteredSorted(ctx context.Context, name, grade, sortBy, order string, limit, offset int) ([]model.Student, error) {
	// Use an allowlist because SQL parameters cannot safely represent column names or keywords.
	columns := map[string]string{"id": "id", "name": "name", "grade": "grade"}
	column, ok := columns[sortBy]
	if !ok {
		column = "id"
	}
	if order != "asc" && order != "desc" {
		order = "asc"
	}
	query := fmt.Sprintf(`SELECT id, name, grade FROM students
  WHERE ($1 = '' OR name ILIKE '%%' || $1 || '%%')
    AND ($2 = '' OR grade ILIKE '%%' || $2 || '%%')
  ORDER BY %s %s LIMIT $3 OFFSET $4`, column, order)
	rows, err := r.DB.Query(ctx, query, name, grade, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	students := make([]model.Student, 0)
	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Grade); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, rows.Err()
}

func (r *PostgresStudentRepository) CountFiltered(ctx context.Context, name, grade string) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM students
  WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
    AND ($2 = '' OR grade ILIKE '%' || $2 || '%')`, name, grade).Scan(&count)
	return count, err
}
