package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	// ISSUE 1: Hardcoded database username and password.
	// Credentials should be loaded from environment variables or a secret manager.
	connString := "postgres://postgres:postgres@localhost:5432/gotraining"

	// ISSUE 2: context.Background() has no timeout.
	// Database operations can wait indefinitely if PostgreSQL becomes unavailable.
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	// ==================================================
	// CREATE
	// ==================================================

	fmt.Println("\n----- CREATE EMPLOYEE -----")

	name := "Rajesh"
	email := "rajesh@gmail.com"
	age := 35
	salary := 50000.00

	// ISSUE 3: Employee input is not validated before inserting.
	// Name, email, age, and salary should be validated first.

	var id int

	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO employees
		 (name, email, age, salary)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		name,
		email,
		age,
		salary,
	).Scan(&id)

	if err != nil {
		// ISSUE 4: Program continues after INSERT failure.
		// If INSERT fails, id may remain 0 and subsequent operations are invalid.
		log.Println("Insert error:", err)
	}

	// ISSUE 5: Success message is printed even when INSERT fails.
	fmt.Println("Employee created successfully")
	fmt.Println("Generated ID:", id)

	// ==================================================
	// READ - ONE EMPLOYEE
	// ==================================================

	fmt.Println("\n----- READ EMPLOYEE -----")

	var employeeID int
	var employeeName string
	var employeeEmail string
	var employeeAge int
	var employeeSalary float64

	err = conn.QueryRow(
		context.Background(),
		`SELECT id, name, email, age, salary
		 FROM employees
		 WHERE id = $1`,
		id,
	).Scan(
		&employeeID,
		&employeeName,
		&employeeEmail,
		&employeeAge,
		&employeeSalary,
	)

	if err != nil {
		log.Println("Read error:", err)
	}

	// ISSUE 6: Employee information is printed even when SELECT fails.
	// The program should stop or handle the error before using these values.
	fmt.Println("ID     :", employeeID)
	fmt.Println("Name   :", employeeName)
	fmt.Println("Email  :", employeeEmail)
	fmt.Println("Age    :", employeeAge)
	fmt.Println("Salary :", employeeSalary)

	// ==================================================
	// READ - ALL EMPLOYEES
	// ==================================================

	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	rows, err := conn.Query(
		context.Background(),
		`SELECT id, name, email, age, salary
		 FROM employees
		 ORDER BY id`,
	)

	if err != nil {
		log.Println("Read all error:", err)
	} else {

		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(
				&employeeID,
				&employeeName,
				&employeeEmail,
				&employeeAge,
				&employeeSalary,
			)

			if err != nil {
				// ISSUE 7: Scan error is logged but execution continues.
				// The program should skip the invalid row or return the error.
				log.Println("Scan error:", err)
			}

			fmt.Println("-----------------------------")
			fmt.Println("ID     :", employeeID)
			fmt.Println("Name   :", employeeName)
			fmt.Println("Email  :", employeeEmail)
			fmt.Println("Age    :", employeeAge)
			fmt.Println("Salary :", employeeSalary)
		}

		// ISSUE 8: rows.Err() is not checked.
		// Errors that occur while iterating through rows may be missed.
	}

	// ==================================================
	// UPDATE
	// ==================================================

	fmt.Println("\n----- UPDATE EMPLOYEE -----")

	newSalary := 60000.00

	result, err := conn.Exec(
		context.Background(),
		`UPDATE employees
		 SET salary = $1
		 WHERE id = $2`,
		newSalary,
		id,
	)

	if err != nil {
		log.Println("Update error:", err)
	}

	// ISSUE 9: result may be invalid after Exec() fails.
	// Calling RowsAffected() without handling the error can cause a panic.
	fmt.Println("Employee updated successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	// ==================================================
	// DELETE
	// ==================================================

	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err = conn.Exec(
		context.Background(),
		`DELETE FROM employees
		 WHERE id = $1`,
		id,
	)

	if err != nil {
		log.Println("Delete error:", err)
	}

	// ISSUE 10: Success message is printed even when DELETE fails.
	// Success should only be reported after confirming the operation succeeded.
	fmt.Println("Employee deleted successfully")
	fmt.Println("Rows affected:", result.RowsAffected())
}
