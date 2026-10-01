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
	// SEARCH BY NAME
	// =====================================

	fmt.Println("\n----- SEARCH BY NAME -----")

	rows, err := conn.Query(
		context.Background(),
		`
		SELECT id,name,email,salary
		FROM employees_q3
		WHERE name ILIKE $1
		`,
		"%Raj%",
	)

	if err != nil {
		log.Println(err)
	} else {

		for rows.Next() {

			var id int
			var name string
			var email string
			var salary float64

			rows.Scan(
				&id,
				&name,
				&email,
				&salary,
			)

			fmt.Println(
				id,
				name,
				email,
				salary,
			)

		}

		rows.Close()

	}

	// =====================================
	// SEARCH BY DEPARTMENT
	// =====================================

	fmt.Println("\n----- SEARCH BY DEPARTMENT -----")

	rows, err = conn.Query(
		context.Background(),
		`
		SELECT 
		e.name,
		e.salary,
		d.name

		FROM employees_q3 e

		JOIN departments_q3 d

		ON e.department_id=d.id

		WHERE d.name=$1
		`,
		"IT",
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var employee string
			var salary float64
			var department string

			rows.Scan(
				&employee,
				&salary,
				&department,
			)

			fmt.Println(
				employee,
				salary,
				department,
			)

		}

		rows.Close()

	}

	// =====================================
	// SEARCH BY SALARY
	// =====================================

	fmt.Println("\n----- SEARCH BY SALARY -----")

	rows, err = conn.Query(
		context.Background(),
		`
		SELECT name,salary
		FROM employees_q3
		WHERE salary >= $1
		`,
		50000,
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var name string
			var salary float64

			rows.Scan(
				&name,
				&salary,
			)

			fmt.Println(
				name,
				salary,
			)

		}

		rows.Close()

	}

	// =====================================
	// PAGINATION
	// LIMIT OFFSET
	// =====================================

	fmt.Println("\n----- PAGINATION -----")

	page := 1
	pageSize := 2

	offset := (page - 1) * pageSize

	rows, err = conn.Query(
		context.Background(),
		`
		SELECT id,name,salary
		FROM employees_q3

		ORDER BY id ASC

		LIMIT $1 OFFSET $2
		`,
		pageSize,
		offset,
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var id int
			var name string
			var salary float64

			rows.Scan(
				&id,
				&name,
				&salary,
			)

			fmt.Println(
				id,
				name,
				salary,
			)

		}

		rows.Close()

	}

	// =====================================
	// SORTING
	// =====================================

	fmt.Println("\n----- SORTING BY SALARY -----")

	rows, err = conn.Query(
		context.Background(),
		`
		SELECT name,salary

		FROM employees_q3

		ORDER BY salary DESC
		`,
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var name string
			var salary float64

			rows.Scan(
				&name,
				&salary,
			)

			fmt.Println(
				name,
				salary,
			)

		}

		rows.Close()

	}

}
