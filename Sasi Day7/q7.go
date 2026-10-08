package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:sasi2356@localhost:5432/assignment_db"

func deposit(ctx context.Context, conn *pgx.Conn, accID int, amount float64) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, accID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO transactions (account_id, txn_type, amount) VALUES ($1, 'DEPOSIT', $2)", accID, amount)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY, holder_name VARCHAR(100), balance NUMERIC(12,2) DEFAULT 0.00);
		CREATE TABLE IF NOT EXISTS transactions (id SERIAL PRIMARY KEY, account_id INT, txn_type VARCHAR(20), amount NUMERIC(12,2), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	`)

	var accID int
	_ = conn.QueryRow(ctx, "INSERT INTO accounts (holder_name, balance) VALUES ('Frank', 1000.00) RETURNING id").Scan(&accID)

	_ = deposit(ctx, conn, accID, 500.00)

	var balance float64
	_ = conn.QueryRow(ctx, "SELECT balance FROM accounts WHERE id = $1", accID).Scan(&balance)
	fmt.Printf("Account ID: %d | New Balance after Deposit: %.2f\n", accID, balance)
}