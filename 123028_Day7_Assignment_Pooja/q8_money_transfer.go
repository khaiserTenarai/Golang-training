package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:admin@localhost:5432/assignment_db"

func transferMoney(ctx context.Context, conn *pgx.Conn, senderID, receiverID int, amount float64) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // Ensures rollback on error

	// Check sender balance
	var senderBal float64
	err = tx.QueryRow(ctx, "SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", senderID).Scan(&senderBal)
	if err != nil {
		return fmt.Errorf("sender account not found: %w", err)
	}
	if senderBal < amount {
		return fmt.Errorf("insufficient balance: current balance is %.2f", senderBal)
	}

	// Deduct from sender
	_, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, senderID)
	if err != nil {
		return err
	}

	// Add to receiver
	_, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx) // Commit transaction
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY, holder_name VARCHAR(100), balance NUMERIC(12,2));
	`)

	var acc1, acc2 int
	_ = conn.QueryRow(ctx, "INSERT INTO accounts (holder_name, balance) VALUES ('Grace', 500.00) RETURNING id").Scan(&acc1)
	_ = conn.QueryRow(ctx, "INSERT INTO accounts (holder_name, balance) VALUES ('Heidi', 200.00) RETURNING id").Scan(&acc2)

	fmt.Printf("Attempting transfer of 300.00 from Account %d to Account %d...\n", acc1, acc2)
	err = transferMoney(ctx, conn, acc1, acc2, 300.00)
	if err != nil {
		log.Fatalf("Transfer failed: %v\n", err)
	}
	fmt.Println("Transfer completed successfully!")
}
