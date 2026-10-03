package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	connString := "postgres://postgres:Info%40131@localhost:5432/go_training"

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
	// CREATE CUSTOMER
	// =====================================

	fmt.Println("\n----- CREATE CUSTOMER -----")

	var customerID int

	err = conn.QueryRow(
		context.Background(),
		`
		INSERT INTO customers
		(name,email,phone,city)

		VALUES($1,$2,$3,$4)

		RETURNING id
		`,
		"Kiran",
		"kiran@gmail.com",
		"9876543214",
		"Pune",
	).Scan(&customerID)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Customer created ID:", customerID)

	}

	// =====================================
	// READ CUSTOMER
	// =====================================

	fmt.Println("\n----- READ CUSTOMER -----")

	var name string
	var email string
	var phone string
	var city string

	err = conn.QueryRow(
		context.Background(),
		`
		SELECT name,email,phone,city

		FROM customers

		WHERE id=$1
		`,
		customerID,
	).Scan(
		&name,
		&email,
		&phone,
		&city,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Phone:", phone)
		fmt.Println("City:", city)

	}

	// =====================================
	// UPDATE CUSTOMER
	// =====================================

	fmt.Println("\n----- UPDATE CUSTOMER -----")

	result, err := conn.Exec(
		context.Background(),
		`
		UPDATE customers

		SET city=$1

		WHERE id=$2
		`,
		"Delhi",
		customerID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Customer updated")
		fmt.Println("Rows:", result.RowsAffected())

	}

	// =====================================
	// SEARCH CUSTOMER
	// =====================================

	fmt.Println("\n----- SEARCH CUSTOMER -----")

	rows, err := conn.Query(
		context.Background(),
		`
		SELECT id,name,email,city

		FROM customers

		WHERE name ILIKE $1
		`,
		"%Ra%",
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var id int
			var cname string
			var cemail string
			var ccity string

			rows.Scan(
				&id,
				&cname,
				&cemail,
				&ccity,
			)

			fmt.Println(
				id,
				cname,
				cemail,
				ccity,
			)

		}

		rows.Close()

	}

	// =====================================
	// PAGINATION
	// =====================================

	fmt.Println("\n----- PAGINATION -----")

	page := 1
	pageSize := 2

	offset := (page - 1) * pageSize

	rows, err = conn.Query(
		context.Background(),
		`
		SELECT id,name,email,city

		FROM customers

		ORDER BY id

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
			var cname string
			var cemail string
			var ccity string

			rows.Scan(
				&id,
				&cname,
				&cemail,
				&ccity,
			)

			fmt.Println(
				id,
				cname,
				cemail,
				ccity,
			)

		}

		rows.Close()

	}

	// =====================================
	// DELETE CUSTOMER
	// =====================================

	fmt.Println("\n----- DELETE CUSTOMER -----")

	result, err = conn.Exec(
		context.Background(),
		`
		DELETE FROM customers

		WHERE id=$1
		`,
		customerID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Customer deleted")
		fmt.Println("Rows:", result.RowsAffected())

	}

}
