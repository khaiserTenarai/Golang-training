package repository
import (
	"context"

	"ems/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

/*
	EmployeeRepositoryImpl implements
	EmployeeRepository.
*/
type EmployeeRepositoryImpl struct {
	db *pgxpool.Pool
}

/*
	Constructor.
*/
func NewEmployeeRepository(
	db *pgxpool.Pool,
) EmployeeRepository {

	return &EmployeeRepositoryImpl{
		db: db,
	}
}

/*
	Save employee into PostgreSQL.
*/
func (r EmployeeRepositoryImpl) Save(
	employee model.Employee,
) error {

	query := `
		INSERT INTO employees
		(name, age, email, salary)
		VALUES ($1, $2, $3, $4)
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

/*
	Find one employee using ID.
*/
func (r EmployeeRepositoryImpl) FindByID(
	id int,
) (model.Employee, error) {

	query := `
		SELECT id, name, age, email, salary
		FROM employees
		WHERE id = $1
	`

	var employee model.Employee

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

	return employee, err
}

/*
	Find all employees.

	This method is especially important
	for Day 9 because the service gets
	all employees and sends them to workers.
*/
func (r EmployeeRepositoryImpl) FindAll() (
	[]model.Employee,
	error,
) {

	query := `
		SELECT id, name, age, email, salary
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

/*
	Update employee.
*/
func (r EmployeeRepositoryImpl) Update(
	employee model.Employee,
) error {

	query := `
		UPDATE employees
		SET name = $1,
		    age = $2,
		    email = $3,
		    salary = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.Salary,
		employee.ID,
	)

	return err
}

/*
	Delete employee.
*/
func (r EmployeeRepositoryImpl) Delete(
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