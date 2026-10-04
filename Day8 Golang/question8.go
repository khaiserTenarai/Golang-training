package main

import (
	"context"
	"database/sql"
	"errors"
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

	ctx := context.Background()

	err = transferMoney(ctx, db, 1, 2, 150.00)
	if err != nil {
		log.Fatalf("Transfer failed: %v", err)
	}

	fmt.Println("Transfer successful")
}

func transferMoney(ctx context.Context, db *sql.DB, fromAccountID, toAccountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	if fromAccountID == toAccountID {
		return errors.New("cannot transfer to the same account")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	firstLock, secondLock := fromAccountID, toAccountID
	if fromAccountID > toAccountID {
		firstLock, secondLock = toAccountID, fromAccountID
	}

	_, err = tx.ExecContext(ctx, `SELECT id FROM accounts WHERE id = $1 FOR UPDATE`, firstLock)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `SELECT id FROM accounts WHERE id = $1 FOR UPDATE`, secondLock)
	if err != nil {
		return err
	}

	var currentBalance float64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id = $1`, fromAccountID).Scan(&currentBalance)
	if err != nil {
		return err
	}

	if currentBalance < amount {
		return errors.New("insufficient funds")
	}

	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, fromAccountID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, toAccountID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO transactions (account_id, type, amount) 
		VALUES ($1, 'transfer_out', $2), ($3, 'transfer_in', $4)`,
		fromAccountID, amount, toAccountID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}