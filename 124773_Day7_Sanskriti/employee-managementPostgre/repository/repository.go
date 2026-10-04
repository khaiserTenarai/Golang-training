package repository

import (
	"context"

	"employee-managementPostgre/model"

	"github.com/jackc/pgx/v5"
)

// EmployeeRepository defines employee database operations.
type EmployeeRepository interface {
	CreateEmployee(employee *model.Employee) error
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployee(employee *model.Employee) error
	DeleteEmployee(id int) error
}

// EmployeeRepositoryImpl implements EmployeeRepository.
type EmployeeRepositoryImpl struct {
	DB *pgx.Conn
}

// CreateEmployee inserts a new employee.
func (r *EmployeeRepositoryImpl) CreateEmployee(employee *model.Employee) error {

	query := `
		INSERT INTO employees
		(name, email, age, salary, department_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.DepartmentID,
	).Scan(&employee.ID)

	return err
}

// GetAllEmployees returns all employees.
func (r *EmployeeRepositoryImpl) GetAllEmployees() ([]model.Employee, error) {

	query := `
		SELECT id, name, email, age, salary, department_id
		FROM employees
		ORDER BY id
	`

	rows, err := r.DB.Query(context.Background(), query)

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
			&employee.DepartmentID,
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

// GetEmployeeByID returns one employee.
func (r *EmployeeRepositoryImpl) GetEmployeeByID(id int) (model.Employee, error) {

	query := `
		SELECT id, name, email, age, salary, department_id
		FROM employees
		WHERE id = $1
	`

	var employee model.Employee

	err := r.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Age,
		&employee.Salary,
		&employee.DepartmentID,
	)

	return employee, err
}

// UpdateEmployee updates an employee.
func (r *EmployeeRepositoryImpl) UpdateEmployee(employee *model.Employee) error {

	query := `
		UPDATE employees
		SET name = $1,
			email = $2,
			age = $3,
			salary = $4,
			department_id = $5
		WHERE id = $6
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.DepartmentID,
		employee.ID,
	)

	return err
}

// DeleteEmployee deletes an employee.
func (r *EmployeeRepositoryImpl) DeleteEmployee(id int) error {

	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}