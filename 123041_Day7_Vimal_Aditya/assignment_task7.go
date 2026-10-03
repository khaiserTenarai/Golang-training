package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	// CREATE ACCOUNT
	// =====================================

	fmt.Println("\n----- CREATE ACCOUNT -----")

	var accountID int

	err = conn.QueryRow(
		context.Background(),
		`
		INSERT INTO accounts
		(account_number, holder_name, balance)

		VALUES($1,$2,$3)

		RETURNING id
		`,
		"100001",
		"Rajesh",
		5000,
	).Scan(&accountID)

	if err != nil {

		log.Println("Account creation error:", err)

	} else {

		fmt.Println("Account created ID:", accountID)

	}

	// =====================================
	// DEPOSIT MONEY
	// =====================================

	fmt.Println("\n----- DEPOSIT -----")

	depositAmount := 3000

	_, err = conn.Exec(
		context.Background(),
		`
		UPDATE accounts

		SET balance = balance + $1

		WHERE id=$2
		`,
		depositAmount,
		accountID,
	)

	if err != nil {

		log.Println("Deposit error:", err)

	} else {

		fmt.Println("Deposit successful")

	}

	// Insert deposit history

	_, err = conn.Exec(
		context.Background(),
		`
		INSERT INTO transactions
		(account_id, transaction_type, amount)

		VALUES($1,$2,$3)
		`,
		accountID,
		"DEPOSIT",
		depositAmount,
	)

	if err != nil {

		log.Println("Transaction history error:", err)

	}

	// =====================================
	// WITHDRAW MONEY
	// =====================================

	fmt.Println("\n----- WITHDRAW -----")

	withdrawAmount := 2000

	result, err := conn.Exec(
		context.Background(),
		`
		UPDATE accounts

		SET balance = balance - $1

		WHERE id=$2

		AND balance >= $1
		`,
		withdrawAmount,
		accountID,
	)

	if err != nil {

		log.Println("Withdrawal error:", err)

	} else {

		if result.RowsAffected() == 0 {

			fmt.Println("Insufficient balance")

		} else {

			fmt.Println("Withdrawal successful")

			// Insert withdrawal history

			_, err = conn.Exec(
				context.Background(),
				`
				INSERT INTO transactions
				(account_id, transaction_type, amount)

				VALUES($1,$2,$3)
				`,
				accountID,
				"WITHDRAW",
				withdrawAmount,
			)

			if err != nil {

				log.Println(err)

			}

		}

	}

	// =====================================
	// BALANCE ENQUIRY
	// =====================================

	fmt.Println("\n----- BALANCE ENQUIRY -----")

	var balance float64

	err = conn.QueryRow(
		context.Background(),
		`
		SELECT balance

		FROM accounts

		WHERE id=$1
		`,
		accountID,
	).Scan(&balance)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Current Balance:", balance)

	}

	// =====================================
	// TRANSACTION HISTORY
	// =====================================

	fmt.Println("\n----- TRANSACTION HISTORY -----")

	rows, err := conn.Query(
		context.Background(),
		`
		SELECT 
		transaction_type,
		amount,
		transaction_date

		FROM transactions

		WHERE account_id=$1

		ORDER BY transaction_date
		`,
		accountID,
	)

	if err != nil {

		log.Println(err)

	} else {

		defer rows.Close()

		for rows.Next() {

			var transactionType string
			var amount float64
			var date time.Time

			err := rows.Scan(
				&transactionType,
				&amount,
				&date,
			)

			if err != nil {

				log.Println("Scan error:", err)
				continue

			}

			fmt.Println("----------------")
			fmt.Println("Type:", transactionType)
			fmt.Println("Amount:", amount)
			fmt.Println(
				"Date:",
				date.Format("2006-01-02 15:04:05"),
			)

		}

	}

}
