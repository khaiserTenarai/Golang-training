package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	connString := "postgres://postgres:admin@localhost:5432/gotraining"

	conn, err := pgx.Connect(
		context.Background(),
		connString,
	)

	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	// =====================================
	// SALARY UPDATE USING TRANSACTION
	// =====================================

	ctx := context.Background()

	tx, err := conn.Begin(ctx)

	if err != nil {

		log.Fatal("Transaction start error:", err)

	}

	defer tx.Rollback(ctx)

	employeeID := 1
	newSalary := 70000.00

	// -------------------------------------
	// GET OLD SALARY
	// -------------------------------------

	var oldSalary float64

	err = tx.QueryRow(
		ctx,
		`
		SELECT salary
		FROM employees_q4
		WHERE id=$1
		`,
		employeeID,
	).Scan(&oldSalary)

	if err != nil {

		log.Println("Employee not found")
		return

	}

	fmt.Println("Old Salary:", oldSalary)

	// -------------------------------------
	// UPDATE EMPLOYEE SALARY
	// -------------------------------------

	_, err = tx.Exec(
		ctx,
		`
		UPDATE employees_q4

		SET salary=$1

		WHERE id=$2
		`,
		newSalary,
		employeeID,
	)

	if err != nil {

		log.Println("Salary update failed")
		return

	}

	// -------------------------------------
	// INSERT SALARY HISTORY
	// -------------------------------------

	_, err = tx.Exec(
		ctx,
		`
		INSERT INTO salary_history_q4
		(employee_id,old_salary,new_salary)

		VALUES($1,$2,$3)
		`,
		employeeID,
		oldSalary,
		newSalary,
	)

	if err != nil {

		log.Println("History insert failed")
		return

	}

	// =====================================
	// COMMIT TRANSACTION
	// =====================================

	err = tx.Commit(ctx)

	if err != nil {

		log.Println("Commit failed")

	} else {

		fmt.Println("Salary updated successfully")

	}

	// =====================================
	// DISPLAY SALARY HISTORY
	// =====================================

	fmt.Println("\n----- SALARY HISTORY -----")

	rows, err := conn.Query(
		ctx,
		`
		SELECT
		employee_id,
		old_salary,
		new_salary,
		updated_at

		FROM salary_history_q4
		`,
	)

	if err != nil {

		log.Println(err)

	} else {

		defer rows.Close()

		for rows.Next() {

			var empID int
			var oldSal float64
			var newSal float64
			var date string

			rows.Scan(
				&empID,
				&oldSal,
				&newSal,
				&date,
			)

			fmt.Println("------------------")
			fmt.Println("Employee ID:", empID)
			fmt.Println("Old Salary:", oldSal)
			fmt.Println("New Salary:", newSal)
			fmt.Println("Updated:", date)

		}

	}

}
