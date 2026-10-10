package repository

import (
 "context"
 "fmt"
 "strings"

 "student-management/model"

 "github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStudentRepository struct { DB *pgxpool.Pool }

func NewPostgresStudentRepository(db *pgxpool.Pool) StudentRepository {
 return &PostgresStudentRepository{DB: db}
}

func (r *PostgresStudentRepository) Create(ctx context.Context, student model.Student) error {
 _, err := r.DB.Exec(ctx, `INSERT INTO students (name, grade) VALUES ($1, $2)`, student.Name, student.Grade)
 return err
}

func (r *PostgresStudentRepository) List(ctx context.Context, name, grade, sortBy, order string, limit, offset int) ([]model.Student, error) {
 // Map query values to known SQL column names before building ORDER BY.
 columns := map[string]string{"id": "id", "name": "name", "grade": "grade"}
 column, ok := columns[strings.ToLower(sortBy)]
 if !ok { column = "id" }
 direction := "ASC"
 if strings.ToLower(order) == "desc" { direction = "DESC" }

 query := fmt.Sprintf(`
  SELECT id, name, grade FROM students
  WHERE ($1 = '' OR name ILIKE '%%' || $1 || '%%')
    AND ($2 = '' OR grade ILIKE '%%' || $2 || '%%')
  ORDER BY %s %s LIMIT $3 OFFSET $4`, column, direction)

 rows, err := r.DB.Query(ctx, query, name, grade, limit, offset)
 if err != nil { return nil, err }
 defer rows.Close()

 students := make([]model.Student, 0)
 for rows.Next() {
  var student model.Student
  if err := rows.Scan(&student.ID, &student.Name, &student.Grade); err != nil { return nil, err }
  students = append(students, student)
 }
 if err := rows.Err(); err != nil { return nil, err }
 return students, nil
}

func (r *PostgresStudentRepository) Count(ctx context.Context, name, grade string) (int, error) {
 var total int
 err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM students WHERE ($1 = '' OR name ILIKE '%' || $1 || '%') AND ($2 = '' OR grade ILIKE '%' || $2 || '%')`, name, grade).Scan(&total)
 return total, err
}
