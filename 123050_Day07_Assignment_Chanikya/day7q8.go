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

	ctx := context.Background()

	// =====================================
	// MONEY TRANSFER DETAILS
	// =====================================

	fromAccount := 1
	toAccount := 2

	// float64 because PostgreSQL NUMERIC is scanned as float64
	transferAmount := 3000.00

	// =====================================
	// START TRANSACTION
	// =====================================

	tx, err := conn.Begin(ctx)

	if err != nil {

		log.Fatal("Transaction start error:", err)

	}

	// Rollback if any error occurs

	defer tx.Rollback(ctx)

	// =====================================
	// CHECK SENDER BALANCE
	// =====================================

	var balance float64

	err = tx.QueryRow(
		ctx,
		`
		SELECT balance

		FROM accounts_q8

		WHERE id=$1

		FOR UPDATE
		`,
		fromAccount,
	).Scan(&balance)

	if err != nil {

		log.Println("Sender account not found")
		return

	}

	fmt.Println("Sender Balance:", balance)

	if balance < transferAmount {

		fmt.Println("Insufficient balance")
		return

	}

	// =====================================
	// DEDUCT MONEY FROM SENDER
	// =====================================

	_, err = tx.Exec(
		ctx,
		`
		UPDATE accounts_q8

		SET balance = balance - $1

		WHERE id=$2
		`,
		transferAmount,
		fromAccount,
	)

	if err != nil {

		log.Println("Debit failed:", err)
		return

	}

	// =====================================
	// ADD MONEY TO RECEIVER
	// =====================================

	_, err = tx.Exec(
		ctx,
		`
		UPDATE accounts_q8

		SET balance = balance + $1

		WHERE id=$2
		`,
		transferAmount,
		toAccount,
	)

	if err != nil {

		log.Println("Credit failed:", err)
		return

	}

	// =====================================
	// COMMIT TRANSACTION
	// =====================================

	err = tx.Commit(ctx)

	if err != nil {

		log.Println("Commit failed:", err)

	} else {

		fmt.Println("Money transferred successfully")

	}

	// =====================================
	// DISPLAY FINAL BALANCE
	// =====================================

	fmt.Println("\n----- ACCOUNT BALANCE -----")

	rows, err := conn.Query(
		ctx,
		`
		SELECT
		id,
		holder_name,
		balance

		FROM accounts_q8

		ORDER BY id
		`,
	)

	if err != nil {

		log.Println(err)

	} else {

		defer rows.Close()

		for rows.Next() {

			var id int
			var name string
			var finalBalance float64

			err := rows.Scan(
				&id,
				&name,
				&finalBalance,
			)

			if err != nil {

				log.Println(err)
				continue

			}

			fmt.Println("----------------")
			fmt.Println("ID:", id)
			fmt.Println("Name:", name)
			fmt.Println("Balance:", finalBalance)

		}

	}

}
