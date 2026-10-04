package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management-app/model"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepositoryImpl struct {
	db *pgxpool.Pool
}

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewEmployeeRepository(
	db *pgxpool.Pool,
) EmployeeRepository {

	return &EmployeeRepositoryImpl{
		db: db,
	}
}

// ==================================================
// CREATE
// ==================================================

func (r *EmployeeRepositoryImpl) Save(
	employee model.Employee,
) error {

	fmt.Println("\n----- CREATE EMPLOYEE -----")

	query := `
		INSERT INTO employees
			(name, email, age, salary, department_id)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.DepartmentID,
	).Scan(&employee.ID)

	if err != nil {
		return err
	}

	fmt.Println("Employee created successfully.")
	fmt.Println("Generated ID:", employee.ID)

	return nil
}

// ==================================================
// READ ONE
// ==================================================

func (r *EmployeeRepositoryImpl) FindByID(
	id int,
) (model.Employee, error) {

	fmt.Println("\n----- READ EMPLOYEE -----")

	var employee model.Employee

	query := `
		SELECT
			e.id,
			e.name,
			e.email,
			e.age,
			e.salary,
			e.department_id,
			d.name
		FROM employees e
		JOIN departments d
			ON e.department_id = d.id
		WHERE e.id = $1
	`

	err := r.db.QueryRow(
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
		&employee.DepartmentName,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return model.Employee{}, ErrNotFound
		}

		return model.Employee{}, err
	}

	return employee, nil
}

// ==================================================
// READ ALL
// ==================================================

func (r *EmployeeRepositoryImpl) FindAll() []model.Employee {

	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	query := `
		SELECT
			e.id,
			e.name,
			e.email,
			e.age,
			e.salary,
			e.department_id,
			d.name
		FROM employees e
		JOIN departments d
			ON e.department_id = d.id
		ORDER BY e.id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		fmt.Println("Error fetching employees:", err)
		return []model.Employee{}
	}

	defer rows.Close()

	employees := make([]model.Employee, 0)

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Age,
			&employee.Salary,
			&employee.DepartmentID,
			&employee.DepartmentName,
		)

		if err != nil {
			fmt.Println("Error scanning employee:", err)
			return []model.Employee{}
		}

		employees = append(
			employees,
			employee,
		)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error reading employees:", err)
		return []model.Employee{}
	}

	return employees
}

// ==================================================
// UPDATE
// ==================================================

func (r *EmployeeRepositoryImpl) Update(
	employee model.Employee,
) error {

	fmt.Println("\n----- UPDATE EMPLOYEE -----")

	query := `
		UPDATE employees
		SET
			name = $1,
			email = $2,
			age = $3,
			salary = $4,
			department_id = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.DepartmentID,
		employee.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	fmt.Println("Employee updated successfully.")
	fmt.Println("Employee ID:", employee.ID)

	return nil
}

// ==================================================
// DELETE
// ==================================================

func (r *EmployeeRepositoryImpl) Delete(
	id int,
) error {

	fmt.Println("\n----- DELETE EMPLOYEE -----")

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
		return ErrNotFound
	}

	fmt.Println("Employee deleted successfully.")
	fmt.Println("Employee ID:", id)

	return nil
}
