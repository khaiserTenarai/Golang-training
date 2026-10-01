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

func NewPostgresEmployeeRepository(
	db *pgxpool.Pool,
) *PostgresEmployeeRepository {

	return &PostgresEmployeeRepository{
		db: db,
	}
}

func (r *PostgresEmployeeRepository) Save(
	employee model.Employee,
) error {

	fmt.Println("\n----- CREATE EMPLOYEE -----")

	query := `
		INSERT INTO employees
		(name, email, age, salary, city, state, pincode)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
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
	).Scan(&employee.ID)

	if err != nil {
		return err
	}

	fmt.Println("Employee created successfully")
	fmt.Println("Generated ID:", employee.ID)

	return nil
}


func (r *PostgresEmployeeRepository) FindByID(
	id int64,
) (model.Employee, error) {

	fmt.Println("\n----- READ EMPLOYEE -----")

	var employee model.Employee

	query := `
		SELECT
			id,
			name,
			email,
			age,
			salary,
			city,
			state,
			pincode
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
		&employee.Age,
		&employee.Salary,
		&employee.Address.City,
		&employee.Address.State,
		&employee.Address.Pincode,
	)

	if err != nil {

		if err == pgx.ErrNoRows {
			return employee, fmt.Errorf("employee not found")
		}

		return employee, err
	}

	fmt.Println("ID      :", employee.ID)
	fmt.Println("Name    :", employee.Name)
	fmt.Println("Email   :", employee.Email)
	fmt.Println("Age     :", employee.Age)
	fmt.Println("Salary  :", employee.Salary)
	fmt.Println("City    :", employee.Address.City)
	fmt.Println("State   :", employee.Address.State)
	fmt.Println("Pincode :", employee.Address.Pincode)

	return employee, nil
}


func (r *PostgresEmployeeRepository) FindAll() (
	[]model.Employee,
	error,
) {

	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	query := `
		SELECT
			id,
			name,
			email,
			age,
			salary,
			city,
			state,
			pincode
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
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *PostgresEmployeeRepository) Update(
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
			city = $5,
			state = $6,
			pincode = $7
		WHERE id = $8
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
		employee.ID,
	)

	if err != nil {
		return err
	}

	fmt.Println("Employee updated successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *PostgresEmployeeRepository) Delete(
	id int64,
) error {

	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err := r.db.Exec(
		context.Background(),
		`DELETE FROM employees WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	fmt.Println("Employee deleted successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *PostgresEmployeeRepository) UpdateSalaryWithHistory(id int64, newSalary float64) error {
    ctx := context.Background()

    tx, err := r.db.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback(ctx)

    var currentSalary float64
    err = tx.QueryRow(ctx, `SELECT salary FROM employees WHERE id = $1 FOR UPDATE`, id).Scan(&currentSalary)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return ErrNotFound
        }
        return fmt.Errorf("failed to fetch current salary: %w", err)
    }

    updateQuery := `UPDATE employees SET salary = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
    _, err = tx.Exec(ctx, updateQuery, newSalary, id)
    if err != nil {
        return fmt.Errorf("failed to update salary: %w", err)
    }

    historyQuery := `
        INSERT INTO salary_history (employee_id, old_salary, new_salary)
        VALUES ($1, $2, $3)
    `
    _, err = tx.Exec(ctx, historyQuery, id, currentSalary, newSalary)
    if err != nil {
        return fmt.Errorf("failed to log salary history: %w", err)
    }

    if err := tx.Commit(ctx); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

func (r *PostgresEmployeeRepository) GetSalaryHistory(employeeID int64) ([]model.SalaryHistory, error) {
    ctx := context.Background()

    query := `
        SELECT id, employee_id, old_salary, new_salary, changed_at
        FROM salary_history
        WHERE employee_id = $1
        ORDER BY changed_at DESC
    `
    rows, err := r.db.Query(ctx, query, employeeID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var history []model.SalaryHistory
    for rows.Next() {
        var h model.SalaryHistory
        if err := rows.Scan(&h.ID, &h.EmployeeID, &h.OldSalary, &h.NewSalary, &h.ChangedAt); err != nil {
            return nil, err
        }
        history = append(history, h)
    }

    return history, nil
}