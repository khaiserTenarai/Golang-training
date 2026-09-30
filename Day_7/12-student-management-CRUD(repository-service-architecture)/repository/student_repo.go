package repository
import (
	"context"
	"errors"

	"student_app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepository interface {
	Save(student model.Student) error
	FindByID(id int) (model.Student, error)
	FindAll() ([]model.Student, error)
	Update(student model.Student) error
	Delete(id int) error
}

type StudentRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewStudentRepository(
	db *pgxpool.Pool,
) StudentRepository {

	return &StudentRepositoryImpl{
		db: db,
	}
}

func (r *StudentRepositoryImpl) Save(
	student model.Student,
) error {

	query := `
		INSERT INTO students
		(name, age, email)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		student.Name,
		student.Age,
		student.Email,
	)

	return err
}

func (r *StudentRepositoryImpl) FindByID(
	id int,
) (model.Student, error) {

	var student model.Student

	query := `
		SELECT id, name, age, email
		FROM students
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&student.ID,
		&student.Name,
		&student.Age,
		&student.Email,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return student, errors.New("student not found")
	}

	return student, err
}

func (r *StudentRepositoryImpl) FindAll() (
	[]model.Student,
	error,
) {

	query := `
		SELECT id, name, age, email
		FROM students
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

	students := make([]model.Student, 0)

	for rows.Next() {

		var student model.Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
			&student.Email,
		)

		if err != nil {
			return nil, err
		}

		students = append(
			students,
			student,
		)
	}

	return students, rows.Err()
}

func (r *StudentRepositoryImpl) Update(
	student model.Student,
) error {

	query := `
		UPDATE students
		SET name = $1,
		    age = $2,
		    email = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		student.Name,
		student.Age,
		student.Email,
		student.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("student not found")
	}

	return nil
}

func (r *StudentRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM students
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
		return errors.New("student not found")
	}

	return nil
}
