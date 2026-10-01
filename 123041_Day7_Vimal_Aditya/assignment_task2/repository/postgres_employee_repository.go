package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"example.com/employee-management/model"
)

var ErrNotFound = errors.New("employee not found")

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) *PostgresEmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}

func (r *PostgresEmployeeRepository) Save(employee model.Employee) error {
	fmt.Println("\n----- CREATE EMPLOYEE -----")

	query := `
		INSERT INTO employees
		(name, email, age, salary, city, state, pincode, department_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.Address.City,
		employee.Address.State,
		employee.Address.Pincode,
		employee.DepartmentID,
	).Scan(&employee.ID)

	if err != nil {
		return err
	}

	fmt.Println("Employee created successfully")
	fmt.Println("Generated ID:", employee.ID)
	return nil
}

func (r *PostgresEmployeeRepository) FindByID(id int64) (model.Employee, error) {
	fmt.Println("\n----- READ EMPLOYEE -----")

	var employee model.Employee

	query := `
		SELECT
			e.id,
			e.name,
			e.email,
			e.age,
			e.salary,
			e.city,
			e.state,
			e.pincode,
			e.department_id,
			COALESCE(d.name, 'Unassigned') AS department_name
		FROM employees e
		LEFT JOIN departments d ON e.department_id = d.id
		WHERE e.id = $1
	`

	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Age,
		&employee.Salary,
		&employee.Address.City,
		&employee.Address.State,
		&employee.Address.Pincode,
		&employee.DepartmentID,
		&employee.DepartmentName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return employee, ErrNotFound
		}
		return employee, err
	}

	employee.Display()
	return employee, nil
}

func (r *PostgresEmployeeRepository) FindAll() ([]model.Employee, error) {
	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	query := `
		SELECT
			e.id,
			e.name,
			e.email,
			e.age,
			e.salary,
			e.city,
			e.state,
			e.pincode,
			e.department_id,
			COALESCE(d.name, 'Unassigned') AS department_name
		FROM employees e
		LEFT JOIN departments d ON e.department_id = d.id
		ORDER BY e.id
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Age,
			&employee.Salary,
			&employee.Address.City,
			&employee.Address.State,
			&employee.Address.Pincode,
			&employee.DepartmentID,
			&employee.DepartmentName,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *PostgresEmployeeRepository) Update(employee model.Employee) error {
	fmt.Println("\n----- UPDATE EMPLOYEE -----")

	query := `
		UPDATE employees
		SET
			name = $1,
			email = $2,
			age = $3,
			salary = $4,
			city = $5,
			state = $6,
			pincode = $7,
			department_id = $8,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $9
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.Address.City,
		employee.Address.State,
		employee.Address.Pincode,
		employee.DepartmentID,
		employee.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	fmt.Println("Employee updated successfully")
	fmt.Println("Rows affected:", result.RowsAffected())
	return nil
}

func (r *PostgresEmployeeRepository) Delete(id int64) error {
	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err := r.db.Exec(context.Background(), `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	fmt.Println("Employee deleted successfully")
	fmt.Println("Rows affected:", result.RowsAffected())
	return nil
}