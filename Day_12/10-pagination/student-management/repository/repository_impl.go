package repository

import (
	"context"

	"student-management/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository handles student database queries.
type PostgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository creates a repository.
func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// GetFiltered returns filtered students for one page.
func (r *PostgresRepository) GetFiltered(
	ctx context.Context,
	name, grade string,
	limit, offset int,
) ([]model.Student, error) {
	query := `
		SELECT id, name, age, grade
		FROM students
		WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
		  AND ($2 = '' OR grade ILIKE '%' || $2 || '%')
		ORDER BY id
		LIMIT $3 OFFSET $4
	`

	rows, err := r.db.Query(
		ctx, query, name, grade, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := make([]model.Student, 0)

	for rows.Next() {
		var student model.Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
			&student.Grade,
		)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

// CountFiltered counts matching students.
func (r *PostgresRepository) CountFiltered(
	ctx context.Context,
	name, grade string,
) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM students
		WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
		  AND ($2 = '' OR grade ILIKE '%' || $2 || '%')
	`

	var total int64

	err := r.db.QueryRow(
		ctx, query, name, grade,
	).Scan(&total)

	return total, err
}

// Create inserts a student into the database.
func (r *PostgresRepository) Create(
	ctx context.Context,
	student model.Student,
) (model.Student, error) {
	query := `
		INSERT INTO students (name, age, grade)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(
		ctx,
		query,
		student.Name,
		student.Age,
		student.Grade,
	).Scan(&student.ID)

	if err != nil {
		return model.Student{}, err
	}

	return student, nil
}
