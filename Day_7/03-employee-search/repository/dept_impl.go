package repository
import (
	"context"
	"errors"

	"employee-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrDepartmentNotFound = errors.New("department not found")

type DepartmentRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewDepartmentRepository(
	db *pgxpool.Pool,
) DepartmentRepository {

	return &DepartmentRepositoryImpl{
		db: db,
	}
}

func (r *DepartmentRepositoryImpl) Save(
	department model.Department,
) error {

	query := `
		INSERT INTO departments (name)
		VALUES ($1)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		department.Name,
	)

	return err
}

func (r *DepartmentRepositoryImpl) FindByID(
	id int,
) (model.Department, error) {

	var department model.Department

	query := `
		SELECT id, name
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

	if err == pgx.ErrNoRows {
		return department, ErrDepartmentNotFound
	}

	return department, err
}

func (r *DepartmentRepositoryImpl) FindAll() (
	[]model.Department,
	error,
) {

	query := `
		SELECT id, name
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

	departments := []model.Department{}

	for rows.Next() {

		var department model.Department

		err := rows.Scan(
			&department.ID,
			&department.Name,
		)

		if err != nil {
			return nil, err
		}

		departments = append(
			departments,
			department,
		)
	}

	return departments, rows.Err()
}

func (r *DepartmentRepositoryImpl) Update(
	department model.Department,
) error {

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

	return nil
}

func (r *DepartmentRepositoryImpl) Delete(
	id int,
) error {

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

	return nil
}
