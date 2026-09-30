package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connString := "postgres://postgres:password@localhost:5432/transferdb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Account")
		fmt.Println("2. Transfer Money")
		fmt.Println("3. Balance Enquiry")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createAccount(ctx, db, reader)
		case "2":
			transferMoney(ctx, db, reader)
		case "3":
			balanceEnquiry(ctx, db, reader)
		case "4":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func createAccount(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account holder name: ")
	name := readLine(reader)

	fmt.Print("Opening balance: ")
	balance, _ := strconv.ParseFloat(readLine(reader), 64)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO accounts (name, balance) VALUES ($1, $2) RETURNING id`, name, balance).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created account with ID:", id)
}

func transferMoney(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("From account ID: ")
	fromID, _ := strconv.Atoi(readLine(reader))

	fmt.Print("To account ID: ")
	toID, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Amount: ")
	amount, _ := strconv.ParseFloat(readLine(reader), 64)

	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer tx.Rollback(ctx)

	var fromBalance float64
	err = tx.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, fromID).Scan(&fromBalance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fromBalance < amount {
		fmt.Println("Insufficient balance, transfer aborted")
		return
	}

	result, err := tx.Exec(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, fromID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Source account not found")
		return
	}

	result, err = tx.Exec(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, toID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Destination account not found")
		return
	}

	_, err = tx.Exec(ctx, `INSERT INTO transactions (account_id, type, amount, related_account_id) VALUES ($1, 'transfer_out', $2, $3)`, fromID, amount, toID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = tx.Exec(ctx, `INSERT INTO transactions (account_id, type, amount, related_account_id) VALUES ($1, 'transfer_in', $2, $3)`, toID, amount, fromID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Transfer successful")
}

func balanceEnquiry(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	var balance float64
	err := db.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1`, id).Scan(&balance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Balance: %.2f\n", balance)
}
