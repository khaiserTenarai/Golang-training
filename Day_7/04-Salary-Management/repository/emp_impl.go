package repository
import (
	"context"
	"errors"

	"salary-management/model"

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
			(name, age, email, salary)
		VALUES
			($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.Salary,
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
			salary
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
		&employee.Salary,
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
			salary
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

	employees := make(
		[]model.Employee,
		0,
	)

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Age,
			&employee.Email,
			&employee.Salary,
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
			email = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("employee not found")
	}

	return nil
}

func (r *EmployeeRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM employees
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
		return errors.New("employee not found")
	}

	return nil
}
