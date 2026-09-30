package repository
import (
	"context"
	"errors"

	"attendance_leave/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindByID(id int) (model.Employee, error)
	FindAll() ([]model.Employee, error)
}

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
		(name, email)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
	)

	return err
}

func (r *EmployeeRepositoryImpl) FindByID(
	id int,
) (model.Employee, error) {

	var employee model.Employee

	query := `
		SELECT id, name, email
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
		&employee.Email,
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
		SELECT id, name, email
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
			&employee.Email,
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
