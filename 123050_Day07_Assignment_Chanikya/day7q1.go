package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	// PostgreSQL connection string
	connString := "postgres://postgres:admin@localhost:5432/gotraining"

	// Connect to PostgreSQL
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

	var id int

	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO employees2
		 (name, email, age, salary)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		name,
		email,
		age,
		salary,
	).Scan(&id)

	if err != nil {
		log.Println("Insert error:", err)
	} else {
		fmt.Println("Employee created successfully")
		fmt.Println("Generated ID:", id)
	}

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
		 FROM employees2
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
	} else {
		fmt.Println("ID     :", employeeID)
		fmt.Println("Name   :", employeeName)
		fmt.Println("Email  :", employeeEmail)
		fmt.Println("Age    :", employeeAge)
		fmt.Println("Salary :", employeeSalary)
	}

	// ==================================================
	// READ - ALL EMPLOYEES
	// ==================================================

	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	rows, err := conn.Query(
		context.Background(),
		`SELECT id, name, email, age, salary
		 FROM employees2
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
				log.Println("Scan error:", err)
				continue
			}

			fmt.Println("-----------------------------")
			fmt.Println("ID     :", employeeID)
			fmt.Println("Name   :", employeeName)
			fmt.Println("Email  :", employeeEmail)
			fmt.Println("Age    :", employeeAge)
			fmt.Println("Salary :", employeeSalary)
		}

		if err := rows.Err(); err != nil {
			log.Println("Rows error:", err)
		}
	}

	// ==================================================
	// UPDATE
	// ==================================================

	fmt.Println("\n----- UPDATE EMPLOYEE -----")

	newSalary := 60000.00

	result, err := conn.Exec(
		context.Background(),
		`UPDATE employees2
		 SET salary = $1
		 WHERE id = $2`,
		newSalary,
		id,
	)

	if err != nil {
		log.Println("Update error:", err)
	} else {
		fmt.Println("Employee updated successfully")
		fmt.Println("Rows affected:", result.RowsAffected())
	}

	// ==================================================
	// DELETE
	// ==================================================

	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err = conn.Exec(
		context.Background(),
		`DELETE FROM employees2
		 WHERE id = $1`,
		id,
	)

	if err != nil {
		log.Println("Delete error:", err)
	} else {
		fmt.Println("Employee deleted successfully")
		fmt.Println("Rows affected:", result.RowsAffected())
	}
}
