package repository

import (
	"context"
	"errors"

	"department-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(
	db *pgxpool.Pool,
) EmployeeRepository {

	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) Save(
	employee model.Employee,
) error {

	query := `
		INSERT INTO employees
		(name, age, email, department_id)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.DepartmentID,
	)

	return err
}

func (r *EmployeeRepositoryImpl) FindByID(
	id int,
) (model.Employee, error) {

	var employee model.Employee

	query := `
		SELECT
			id,
			name,
			age,
			email,
			department_id
		FROM employees
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Age,
		&employee.Email,
		&employee.DepartmentID,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return employee, errors.New("employee not found")
	}

	return employee, err
}

func (r *EmployeeRepositoryImpl) FindAll() (
	[]model.Employee,
	error,
) {

	query := `
		SELECT
			id,
			name,
			age,
			email,
			department_id
		FROM employees
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

	employees := make([]model.Employee, 0)

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Age,
			&employee.Email,
			&employee.DepartmentID,
		)

		if err != nil {
			return nil, err
		}

		employees = append(
			employees,
			employee,
		)
	}

	return employees, rows.Err()
}

func (r *EmployeeRepositoryImpl) Update(
	employee model.Employee,
) error {

	query := `
		UPDATE employees
		SET
			name = $1,
			age = $2,
			email = $3,
			department_id = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.DepartmentID,
		employee.ID,
	)

	return err
}

func (r *EmployeeRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}
