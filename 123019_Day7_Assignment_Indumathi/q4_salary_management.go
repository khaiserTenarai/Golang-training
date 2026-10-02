package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:Indu%40123@localhost:5432/assignment_db"

func updateSalaryWithHistory(ctx context.Context, conn *pgx.Conn, empID int, newSalary float64) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Get current salary
	var oldSalary float64
	err = tx.QueryRow(ctx, "SELECT salary FROM employees WHERE id = $1 FOR UPDATE", empID).Scan(&oldSalary)
	if err != nil {
		return fmt.Errorf("employee not found: %w", err)
	}

	// Update employee salary
	_, err = tx.Exec(ctx, "UPDATE employees SET salary = $1 WHERE id = $2", newSalary, empID)
	if err != nil {
		return err
	}

	// Record in history table
	_, err = tx.Exec(ctx, "INSERT INTO salary_history (employee_id, old_salary, new_salary) VALUES ($1, $2, $3)", empID, oldSalary, newSalary)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS employees (id SERIAL PRIMARY KEY, name VARCHAR(100), email VARCHAR(100) UNIQUE, salary NUMERIC(10,2));
		CREATE TABLE IF NOT EXISTS salary_history (id SERIAL PRIMARY KEY, employee_id INT, old_salary NUMERIC(10,2), new_salary NUMERIC(10,2), changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	`)

	var empID int
	_ = conn.QueryRow(ctx, "INSERT INTO employees (name, email, salary) VALUES ('David', 'david@test.com', 50000) ON CONFLICT (email) DO UPDATE SET salary=employees.salary RETURNING id").Scan(&empID)

	err = updateSalaryWithHistory(ctx, conn, empID, 58000.00)
	if err != nil {
		log.Fatalf("Salary update failed: %v\n", err)
	}

	fmt.Println("Salary updated and recorded in history successfully!")
}