package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Account struct {
	ID            int
	AccountNumber string
	OwnerName     string
	Balance       float64
	CreatedAt     time.Time
}

type Transaction struct {
	ID        int
	AccountID int
	Type      string
	Amount    float64
	CreatedAt time.Time
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()

	accID := createAccount(db, "100100100", "Jane Doe")
	
	balance := getBalance(db, accID)
	fmt.Printf("Initial Balance: $%.2f\n", balance)

	err = deposit(ctx, db, accID, 500.00)
	if err != nil {
		log.Fatal(err)
	}

	err = withdraw(ctx, db, accID, 150.00)
	if err != nil {
		log.Fatal(err)
	}

	balance = getBalance(db, accID)
	fmt.Printf("Current Balance: $%.2f\n", balance)

	history := getTransactionHistory(db, accID)
	for _, t := range history {
		fmt.Printf("%s: $%.2f at %v\n", t.Type, t.Amount, t.CreatedAt.Format(time.RFC3339))
	}
}

func createAccount(db *sql.DB, accountNumber, ownerName string) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO accounts (account_number, owner_name, balance) 
		VALUES ($1, $2, 0.00) RETURNING id`,
		accountNumber, ownerName).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func deposit(ctx context.Context, db *sql.DB, accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, accountID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO transactions (account_id, type, amount) 
		VALUES ($1, 'deposit', $2)`, accountID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func withdraw(ctx context.Context, db *sql.DB, accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, accountID).Scan(&currentBalance)
	if err != nil {
		return err
	}

	if currentBalance < amount {
		return errors.New("insufficient funds")
	}

	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, accountID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO transactions (account_id, type, amount) 
		VALUES ($1, 'withdrawal', $2)`, accountID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func getBalance(db *sql.DB, accountID int) float64 {
	var balance float64
	err := db.QueryRow(`SELECT balance FROM accounts WHERE id = $1`, accountID).Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func getTransactionHistory(db *sql.DB, accountID int) []Transaction {
	rows, err := db.Query(`
		SELECT id, account_id, type, amount, created_at 
		FROM transactions 
		WHERE account_id = $1 
		ORDER BY created_at DESC`, accountID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var history []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.AccountID, &t.Type, &t.Amount, &t.CreatedAt); err != nil {
			log.Fatal(err)
		}
		history = append(history, t)
	}
	
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	
	return history
}