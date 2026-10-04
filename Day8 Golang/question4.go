package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = updateEmployeeSalary(context.Background(), db, 1, 110000.00)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Salary updated successfully")
}

func updateEmployeeSalary(ctx context.Context, db *sql.DB, employeeID int, newSalary float64) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var oldSalary float64
	err = tx.QueryRowContext(ctx, `SELECT salary FROM employees WHERE id = $1 FOR UPDATE`, employeeID).Scan(&oldSalary)
	if err != nil {
		return err
	}

	if oldSalary == newSalary {
		return nil
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO salary_history (employee_id, old_salary, new_salary) 
		VALUES ($1, $2, $3)`,
		employeeID, oldSalary, newSalary)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE employees SET salary = $1 WHERE id = $2`, newSalary, employeeID)
	if err != nil {
		return err
	}

	return tx.Commit()
}