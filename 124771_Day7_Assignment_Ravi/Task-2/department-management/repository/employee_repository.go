package repository

import (
	"context"
	"fmt"

	"department-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	DB *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{
		DB: db,
	}
}

func (r *EmployeeRepository) Create(
	ctx context.Context,
	employee *model.Employee,
) error {

	query := `
		INSERT INTO employees
			(name, email, salary, department_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		employee.Name,
		employee.Email,
		employee.Salary,
		employee.DepartmentID,
	).Scan(&employee.ID)

	if err != nil {
		return fmt.Errorf(
			"failed to create employee: %w",
			err,
		)
	}

	return nil
}

func (r *EmployeeRepository) GetByID(
	ctx context.Context,
	id int,
) (*model.EmployeeWithDepartment, error) {

	query := `
		SELECT
			e.id,
			e.name,
			e.email,
			e.salary,
			e.department_id,
			d.name
		FROM employees e
		INNER JOIN departments d
			ON e.department_id = d.id
		WHERE e.id = $1
	`

	var employee model.EmployeeWithDepartment

	err := r.DB.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Salary,
		&employee.DepartmentID,
		&employee.DepartmentName,
	)

	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("employee not found")
		}

		return nil, fmt.Errorf(
			"failed to get employee: %w",
			err,
		)
	}

	return &employee, nil
}

func (r *EmployeeRepository) GetAll(
	ctx context.Context,
) ([]model.EmployeeWithDepartment, error) {

	query := `
		SELECT
			e.id,
			e.name,
			e.email,
			e.salary,
			e.department_id,
			d.name
		FROM employees e
		INNER JOIN departments d
			ON e.department_id = d.id
		ORDER BY e.id
	`

	rows, err := r.DB.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to get employees: %w",
			err,
		)
	}

	defer rows.Close()

	employees := make(
		[]model.EmployeeWithDepartment,
		0,
	)

	for rows.Next() {

		var employee model.EmployeeWithDepartment

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Salary,
			&employee.DepartmentID,
			&employee.DepartmentName,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to scan employee: %w",
				err,
			)
		}

		employees = append(
			employees,
			employee,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"failed while reading employees: %w",
			err,
		)
	}

	return employees, nil
}

func (r *EmployeeRepository) Update(
	ctx context.Context,
	employee *model.Employee,
) error {

	query := `
		UPDATE employees
		SET
			name = $1,
			email = $2,
			salary = $3,
			department_id = $4
		WHERE id = $5
	`

	result, err := r.DB.Exec(
		ctx,
		query,
		employee.Name,
		employee.Email,
		employee.Salary,
		employee.DepartmentID,
		employee.ID,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to update employee: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	return nil
}

func (r *EmployeeRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	result, err := r.DB.Exec(
		ctx,
		query,
		id,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to delete employee: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	return nil
}

func (r *EmployeeRepository) ChangeDepartment(
	ctx context.Context,
	employeeID int,
	departmentID int,
) error {

	query := `
		UPDATE employees
		SET department_id = $1
		WHERE id = $2
	`

	result, err := r.DB.Exec(
		ctx,
		query,
		departmentID,
		employeeID,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to change department: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	return nil
}
